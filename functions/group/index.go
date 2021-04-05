package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/thiduzz/slack-bot/services"
)

func main() {
	groupService := services.GroupService{}
	lambda.Start(groupService.Index)
}