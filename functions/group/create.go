package main

import (
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
func HandleRequest(e InEvent) (OutEvent, error) {
	return OutEvent{
		Payload: e.Payload,
		Status:  200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}