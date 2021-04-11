package services

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/thiduzz/slack-bot/helpers"
	"github.com/thiduzz/slack-bot/repositories"
	"text/template"
)

type MembershipService struct {
	repositories.MembershipRepository
	repositories.GroupRepository
}

func (g MembershipService) Create(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	params := helpers.DecodeRequest(request.Body)
	title := params.Get("text")
	channel := params.Get("channel_id")
	group, err := g.GroupRepository.FindByNameAndChannel(title, channel)
	if err != nil {
		return helpers.NewErrorResponse(err), nil
	}
	t := template.Must(template.New("html-tmpl").ParseFiles([]string{
		"../templates/membership_create.tmpl",
	}...))
	var tpl bytes.Buffer
	err = t.Execute(&tpl, group)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	body, _ := json.Marshal(tpl.String())
	json.HTMLEscape(&buf, body)
	return events.APIGatewayProxyResponse{
		Body: buf.String(),
		IsBase64Encoded: false,
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":           "application/json",
		}},nil
}