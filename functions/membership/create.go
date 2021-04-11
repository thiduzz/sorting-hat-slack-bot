package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/thiduzz/slack-bot/repositories"
	"github.com/thiduzz/slack-bot/services"
)

func main() {
	sess := session.Must(session.NewSession())
	db := dynamodb.New(sess)
	membershipService := services.MembershipService{
		MembershipRepository: repositories.NewMembershipRepository(db),
		GroupRepository:      repositories.NewGroupRepository(db),
	}
	lambda.Start(membershipService.Create)
}