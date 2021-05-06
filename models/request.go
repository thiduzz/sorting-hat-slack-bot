package models

import "github.com/aws/aws-lambda-go/events"

type ProxyRoute struct {
	Entity						string											`json:"proxyEntity,omitempty"`
	Action						string											`json:"proxyAction,omitempty"`

}

type Request struct {
	events.APIGatewayProxyRequest
	DecodedBody                 DecodedBody                          			`json:"decodedBody,omitempty"`
	ProxyRoute					*ProxyRoute										`json:"proxyRoute,omitempty"`
	TriggerId					string											`json:"triggerId,omitempty"`
}

type DecodedBody struct {
	TriggerId	string			`json:"trigger_id,omitempty"`
	Type		string			`json:"type,omitempty"`
	View 		SlackView  		`json:"view,omitempty"`
	Team 		SlackTeam  		`json:"team,omitempty"`
	User		SlackUser		`json:"user,omitempty"`
	State		SlackState		`json:"state,omitempty"`
}

type SlackView struct {
	PrivateMetadata string           `json:"private_metadata,omitempty"`
	CallbackId      string           `json:"callback_id,omitempty"`
}


type SlackTeam struct {
	Id 			string           `json:"id,omitempty"`
	Domain      string           `json:"domain,omitempty"`
}

type SlackUser struct {
	Id 			string           `json:"id,omitempty"`
	Name      string           `json:"name,omitempty"`
	Username      string           `json:"username,omitempty"`
	TeamId      string           `json:"team_id,omitempty"`
}

type SlackState struct {}
