package repositories

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/thiduzz/slack-bot/models"
	"log"
)

type GroupRepository interface {
	Store(group models.Group) error
	IndexByContextReference(channelId string) ([]models.Group, error)
	FindByNameAndChannel(groupName string, channelId string) (*models.Group, error)
	Destroy(group *models.Group) error
}

var _ GroupRepository = &groupDynamo{}

func NewGroupRepository(db *dynamodb.DynamoDB) *groupDynamo {
	return &groupDynamo{db: db}
}

type groupDynamo struct {
	db *dynamodb.DynamoDB
	BaseRepository
}

type GroupListItem struct {
	GroupId   string
	ChannelId string
	Title     string
}

type GroupOwnershipKey struct {
	ContextReference 	string
	GroupId   			string
}

func (g *groupDynamo) Store(group models.Group) error {

	av, err := dynamodbattribute.MarshalMap(group)
	if err != nil {
		return err
	}

	log.Println(fmt.Sprintf("Storing... : %v",av))
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(models.GroupsTableName),
	}
	_, err = g.db.PutItem(input)
	if err != nil {
		return errors.New(fmt.Sprintf("Got error calling PutItem: %s", err))
	}
	return nil
}

func (g *groupDynamo) IndexByContextReference(contextReference string) ([]models.Group, error) {
	filt := expression.Name("context_reference").Equal(expression.Value(contextReference))
	expr, _ := expression.NewBuilder().WithFilter(filt).Build()
	result, err := g.db.Scan(&dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String(models.GroupsTableName),
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Got error calling GetItem: %s", err))
	}
	var groups []models.Group

	for _, i := range result.Items {
		item := models.Group{}
		err = dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Got error unmarshalling: %s", err))
		}
		groups = append(groups, item)
	}
	return groups, nil
}

func (g *groupDynamo) FindByNameAndChannel(groupName string, channelId string) (*models.Group, error) {
	filter := expression.Name("ChannelId").Equal(expression.Value(channelId)).And(expression.Name("Title").Equal(expression.Value(groupName)))
	expr, _ := expression.NewBuilder().WithFilter(filter).Build()
	result, err := g.db.Scan(&dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String(models.GroupsTableName),
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Got error calling Scan: %s", err))
	}
	if len(result.Items) < 0 {
		return nil, errors.New("No group with this name was found.")
	}

	group := models.Group{}
	firstItem := result.Items[0]
	err = dynamodbattribute.UnmarshalMap(firstItem, &group)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Got error unmarshalling: %s", err))
	}
	return &group, nil
}

func (g *groupDynamo) Destroy(group *models.Group) error {

	key, _ := dynamodbattribute.MarshalMap(GroupOwnershipKey{
		GroupId:   			group.GroupId,
		ContextReference: 	group.ContextReference,
	})

	_, err := g.db.DeleteItem(&dynamodb.DeleteItemInput{
		Key:       key,
		TableName: aws.String(models.GroupsTableName),
	})
	if err != nil {
		return errors.New(fmt.Sprintf("Could not delete the group %s : %s", group.Title, err.Error()))
	}
	return nil
}
