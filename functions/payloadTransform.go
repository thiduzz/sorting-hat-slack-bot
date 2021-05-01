package main

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
)


// Event defines your lambda input and output data structure,
// and of course you can have different input and output data structure
type Event struct {
	Payload string `json:"payload"`
	Action string `json:"action"`
	TaskToken string `json:"taskToken"`
}

// HandleRequest handles the incomming StepFunction request
func HandleRequest(e interface{}) (Event, error) {
	var buf bytes.Buffer
	body, _ := json.Marshal(e)
	json.HTMLEscape(&buf, body)
	return Event{
		buf.String(),
		"createGroup",
			"taskTokenExampleThiago",
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}