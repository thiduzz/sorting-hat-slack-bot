package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

// Request defines your lambda input data structure,
type Request struct {
	Body string `json:"body"`
}

// OutEvent defines your lambda output data structure,
type OutEvent struct {
	Response string `json:"response"`
	Status   int    `json:"status"`
}

// HandleRequest handles the incomming StepFunction request
func HandleRequest(e Request) (OutEvent, error) {
	return OutEvent{
		Response: e.Body,
		Status:   200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
