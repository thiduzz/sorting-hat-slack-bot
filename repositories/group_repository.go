package repositories

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/thiduzz/slack-bot/models"
)


type groupRepository struct {
	db *dynamodb.DynamoDB
}

func NewGroupRepository() *groupRepository  {
	sess := session.Must(session.NewSession())
	return &groupRepository{
		db: dynamodb.New(sess),
	}
}

func (g *groupRepository) Store(group *models.Group)  {

}