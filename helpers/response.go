package helpers

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

func NewErrorResponse(err error) events.APIGatewayProxyResponse {
	var buf bytes.Buffer
	body, _ := json.Marshal(map[string]interface{}{
		"response_type": "ephemeral",
		"text":          err.Error(),
	})
	json.HTMLEscape(&buf, body)
	return events.APIGatewayProxyResponse{
		Body:            buf.String(),
		IsBase64Encoded: false,
		StatusCode:      200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		}}
}


func NewInteractivityErrorResponse(errors map[string] string) events.APIGatewayProxyResponse {
	var buf bytes.Buffer
	body, _ := json.Marshal(map[string]interface{}{
		"response_action": 	"errors",
		"errors":          	errors,
	})
	json.HTMLEscape(&buf, body)
	bodyEscaped := buf.String()
	return events.APIGatewayProxyResponse{
		Body: bodyEscaped,
		IsBase64Encoded: false,
		StatusCode:      200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		}}
}
