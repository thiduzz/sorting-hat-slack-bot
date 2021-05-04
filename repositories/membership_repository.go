package repositories

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type MembershipRepository interface {
	Subscribe(userId string) error
}

var _ MembershipRepository = &membershipDynamo{}

func NewMembershipRepository(db *dynamodb.DynamoDB) *membershipDynamo {
	return &membershipDynamo{db: db}
}

type membershipDynamo struct {
	db *dynamodb.DynamoDB
	BaseRepository
}

func (m membershipDynamo) Subscribe(userId string) error {
	panic("implement me")
}
