package helpers

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/thiduzz/slack-bot/repositories"
	"text/template"
)

//FormatListBlockResponse
func FormatListBlockResponse(groups []repositories.GroupListItem) (string, error) {

	var err error
	t, err := template.ParseFiles("../templates/groups_list.tmpl")
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, groups); err != nil {
		return "", err
	}

	return tpl.String(), nil
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