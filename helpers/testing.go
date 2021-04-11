package helpers

import (
	"encoding/base64"
	"github.com/thiduzz/slack-bot/models"
	"net/url"
)

func EncodeToBase64URL(request map[string] string) string {

	params := url.Values{}
	for key, value := range request {
		params.Add(key,value)
	}
	query := params.Encode()
	return base64.StdEncoding.EncodeToString([]byte(query))
}

func GenerateBaseRequest() map[string] string{
	return map[string]string{
		"token": "O8mkcDKXfmIitPp7RXSX4S1U",
		"team_id": "T01T72BF15Z",
		"team_domain": "test",
		"channel_id": "C01T72BFMFV",
		"channel_name": "general",
		"user_id": "U01T02LM6DU",
		"user_name": "test",
		"command": "sorting-hat-group-create",
		"text": "Group Test",
		"api_app_id": "TestingAppId",
		"response_url": "https://hooks.slack.com/command/blablabla",
	}
}

func GenerateBaseGroup() models.Group{
	return models.Group{
		GroupId:     "12345",
		Title:       "Group Test",
		ChannelName: "general",
		ChannelId:   "C01T72BFMFV",
		WorkspaceId: "T01T72BF15Z",
		CreatorId:   "U01T02LM6DU",
		CreatedAt:   "",
	}
}