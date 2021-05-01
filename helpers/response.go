package helpers

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/thiduzz/slack-bot/repositories"
	"io/fs"
	"text/template"
)


func FormatListBlockResponse(fs fs.FS, groups []repositories.GroupListItem) (string, error) {

	var err error
	t, err := template.ParseFS(fs, "templates/list.tmpl")
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