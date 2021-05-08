package models

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"net/url"
)

const (
	CALLBACK_GROUP_STORE = "group.store"
	CALLBACK_GROUP_DELETE = "group.delete"
)

type ProxyRoute struct {
	Entity		string				`json:"proxyEntity,omitempty"`
	Action		string				`json:"proxyAction,omitempty"`

}

type InteractivityRequest struct {
	events.APIGatewayProxyRequest
	DecodedInteractiveBody
	*ProxyRoute
}

type SlashRequest struct {
	events.APIGatewayProxyRequest
	DecodedSlashBody
}

type DecodedInteractiveBody struct {
	TriggerId	string				`json:"trigger_id,omitempty"`
	Type		string				`json:"type,omitempty"`
	View 		SlackView  			`json:"view,omitempty"`
	Team 		SlackTeam  			`json:"team,omitempty"`
	User		SlackUser			`json:"user,omitempty"`
}

type DecodedSlashBody struct {
	ChannelId		string			`schema:"channel_id,omitempty,required,alphanum"`
	WorkspaceId		string			`schema:"team_id,omitempty,required,alphanum"`
	TriggerId		string			`schema:"trigger_id,omitempty,required"`
	UserId 			string   		`schema:"user_id,omitempty,required"`
	Text 			string   		`schema:"text,omitempty"`
	Command 		string   		`schema:"command,omitempty"`
	Token 			string   		`schema:"token,omitempty"`
	ApiAppId 		string   		`schema:"api_app_id,omitempty"`
	ChannelName		string   		`schema:"channel_name,omitempty"`
	UserName		string   		`schema:"user_name,omitempty"`
	ResponseUrl		string   		`schema:"response_url,omitempty"`
	TeamDomain		string   		`schema:"team_domain,omitempty"`
}

type SlackView struct {
	PrivateMetadata string           `json:"private_metadata,omitempty"`
	CallbackId      string           `json:"callback_id,omitempty"`
	State		SlackState			 `json:"state,omitempty"`
}


type SlackTeam struct {
	Id 			string           	`json:"id,omitempty"`
	Domain      string           	`json:"domain,omitempty"`
}

type SlackUser struct {
	Id 			string           	`json:"id,omitempty"`
	Name      	string           	`json:"name,omitempty"`
	Username    string         		`json:"username,omitempty"`
	TeamId      string           	`json:"team_id,omitempty"`
}

type SlackState struct {
	Values	map[string]SlackStateValuesWrapper `json:"values,omitempty"`
}

type SlackStateValuesWrapper map[string] SlackStateValue

type SlackStateValue struct {
	Type 		string           	`json:"type,omitempty"`
	Value      	string           	`json:"value,omitempty"`
}

func NewSlackRequest(payload map[string]interface{}) []byte {
	values, _ := url.ParseQuery("")
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(payload)
	values.Add("payload", reqBodyBytes.String())
	reqBodyBytes = new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(events.APIGatewayProxyRequest{
		Body:                            base64.StdEncoding.EncodeToString([]byte(values.Encode())),
		IsBase64Encoded:                 true,
	})
	return reqBodyBytes.Bytes()
}

func NewSlackSlashRequest(payload map[string]interface{}) []byte {
	values, _ := url.ParseQuery("")
	for s, i := range payload {
		values.Add(s, fmt.Sprintf("%s",i))
	}
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(events.APIGatewayProxyRequest{
		Body:                            base64.StdEncoding.EncodeToString([]byte(values.Encode())),
		IsBase64Encoded:                 true,
	})
	return reqBodyBytes.Bytes()
}