package main

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
)


// Event defines your lambda input and output data structure,
// and of course you can have different input and output data structure
type Event struct {
	Request Request `json:"request"`
	Action string `json:"action"`
}

type Request struct {
	Body interface{} `json:"body"`
}

// HandleRequest handles the incomming StepFunction request
func HandleRequest(e interface{}) (Event, error) {
	var buf bytes.Buffer
	body, _ := json.Marshal(e)
	json.HTMLEscape(&buf, body)
	request := Request{Body: buf.String()}
	return Event{
		request,
		"createGroup",
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}