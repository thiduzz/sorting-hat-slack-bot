package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/thiduzz/slack-bot/models"
	"log"
	"net/url"
)

// Handler function Using AWS Lambda Proxy Request
func HandleCreate(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Payload:")
	log.Println(request.Body)
	sDec, _ := base64.StdEncoding.DecodeString(request.Body)
	params, _ := url.ParseQuery(string(sDec))
	group := models.Group{
		ChannelName: 	params.Get("channel_name"),
		CreatorId:      params.Get("user_id"),
		Title:        	params.Get("text"),
		WorkspaceId:	params.Get("team_id"),
	}
	_, err := govalidator.ValidateStruct(group)
	if err != nil {
		var buf bytes.Buffer
		body, _ := json.Marshal(map[string]interface{}{
			"response_type": "ephemeral",
			"text": err.Error(),
		})
		json.HTMLEscape(&buf, body)
		return events.APIGatewayProxyResponse{
			Body: buf.String(),
			IsBase64Encoded: false,
			StatusCode: 200,
			Headers: map[string]string{
				"Content-Type":           "application/json",
			}}, nil
	}
	return events.APIGatewayProxyResponse{Body: fmt.Sprintf("New group %s created!", group.Title), StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleCreate)
}