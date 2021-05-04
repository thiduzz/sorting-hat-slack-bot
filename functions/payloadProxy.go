package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/thiduzz/slack-bot/helpers"
	"log"
)

func main() {
	//sess := session.Must(session.NewSession())
	//db := dynamodb.New(sess)
	//groupService := services.NewGroupService(db, services.NewSlackService(os.Getenv("SLACK_ACCESS_TOKEN")))
	lambda.Start(Index)
}

func Index(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	params := helpers.DecodeRequest(request.Body)
	log.Println(params)
	return events.APIGatewayProxyResponse{Body: "Ok Thiago!", StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		}}, nil
}
