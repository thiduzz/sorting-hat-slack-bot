package repositories

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/thiduzz/slack-bot/models"
	"log"
)


type groupRepository struct {
	db *dynamodb.DynamoDB
}

type groupListItem struct{
	GroupId string
	ChannelId string
	Title string
}

func NewGroupRepository() *groupRepository  {
	sess := session.Must(session.NewSession())
	return &groupRepository{
		db: dynamodb.New(sess),
	}
}

func (g *groupRepository) Store(group models.Group) error {

	av, err := dynamodbattribute.MarshalMap(group)
	if err != nil {
		return err
	}

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

func (g *groupRepository) IndexByChannelId(channelId string) ([]groupListItem, error) {
	filt := expression.Name("ChannelId").Equal(expression.Value(channelId))
	proj := expression.NamesList(expression.Name("ChannelId"), expression.Name("GroupId"), expression.Name("Title"))
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		log.Fatalf("Got error building expression: %s", err)
	}
	result, err := g.db.Scan(&dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(models.GroupsTableName),
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Got error calling GetItem: %s", err))
	}
	var groups []groupListItem

	for _, i := range result.Items {
		item := groupListItem{}
		err = dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Got error unmarshalling: %s", err))
		}
		groups = append(groups, item)
	}
	return groups, nil
}