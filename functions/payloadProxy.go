package main

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/thiduzz/slack-bot/helpers"
	"github.com/thiduzz/slack-bot/middlewares"
	"github.com/thiduzz/slack-bot/models"
	"github.com/thiduzz/slack-bot/services"
	"os"
)

func main() {
	lambda.Start(
		middlewares.MiddlewareFunc(
			middlewares.ParseRequest(Proxy),
		),
	)
}

func Proxy(ctx context.Context, req *models.Request) (events.APIGatewayProxyResponse, error) {
	switch req.ProxyRoute.Entity {
	case "group":
		sess := session.Must(session.NewSession())
		db := dynamodb.New(sess)
		groupService := services.NewGroupService(db, services.NewSlackService(os.Getenv("SLACK_ACCESS_TOKEN")))
		switch req.ProxyRoute.Action {
		case "store":
			return groupService.Store(ctx, req)
		}
	}
	return helpers.NewErrorResponse(errors.New("not_found")), nil
}