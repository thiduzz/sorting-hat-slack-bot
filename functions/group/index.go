package main

import (
	"embed"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/thiduzz/slack-bot/services"
)

//go:embed templates/*
var static embed.FS

func main() {
	sess := session.Must(session.NewSession())
	db := dynamodb.New(sess)
	groupService := services.NewGroupService(db, static)
	lambda.Start(groupService.Index)
}