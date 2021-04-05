package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

//FormatListBlockResponse
func FormatListBlockResponse(listToFormat []string) string {
	if len(listToFormat) <= 0 {
		return ""
	}
	var stringList string
	for _, item := range listToFormat {
		if listToFormat[len(listToFormat)-1] != item {
			stringList += fmt.Sprintf("• %s\n", item)
		}else{
			stringList += fmt.Sprintf("• %s", item)
		}
	}
	body, _ := json.Marshal(map[string]interface{}{
		"blocks": []interface{}{
			struct{
				TypeName string `json:"type"`
				Text map[string]string `json:"text"`
			}{
					TypeName: "section",
					Text: map[string]string{
						"type": "mrkdwn",
						"text": stringList,
					},
			},
		},
	})
	return string(body)
}

func NewErrorResponse(err error) events.APIGatewayProxyResponse {
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
		}}
}