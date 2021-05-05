package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/thiduzz/slack-bot/middlewares"
	"github.com/thiduzz/slack-bot/services"
	"os"
)

func main() {
	sess := session.Must(session.NewSession())
	db := dynamodb.New(sess)
	groupService := services.NewGroupService(db, services.NewSlackService(os.Getenv("SLACK_ACCESS_TOKEN")))
	lambda.Start(
		middlewares.MiddlewareFunc(
			middlewares.ParseRequest(groupService.Store),
		),
	)
}
