package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)


// InEvent defines your lambda input data structure,
type InEvent struct {
	Payload string `json:"payload"`
	Action  string    `json:"action"`
}

// OutEvent defines your lambda output data structure,
type OutEvent struct {
	Payload string `json:"payload"`
	Status  int    `json:"status"`
}

// HandleRequest handles the incomming StepFunction request
func HandleRequest(e InEvent) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{Body: e.Payload, StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":           "application/json",
		}}, nil
}

func main() {
	lambda.Start(HandleRequest)
}