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
	"github.com/thiduzz/slack-bot/repositories"
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

func Proxy(ctx context.Context, req *models.InteractivityRequest) (events.APIGatewayProxyResponse, error) {

	switch req.Entity {
	case "group":
		sess := session.Must(session.NewSession())
		db := dynamodb.New(sess)
		groupService := services.NewGroupService(db, services.NewSlackService(os.Getenv("SLACK_ACCESS_TOKEN")))
		switch req.Action {
		case "store":
			return groupService.Store(ctx, req)
		}
	case "member":
		sess := session.Must(session.NewSession())
		db := dynamodb.New(sess)
		membershipService := services.MembershipService{
			MembershipRepository: repositories.NewMembershipRepository(db),
			GroupRepository:      repositories.NewGroupRepository(db),
		}
		switch req.Action {
		case "store":
			return membershipService.Store(ctx, req)
		}

	}
	return helpers.NewErrorResponse(errors.New("not_found")), nil
}