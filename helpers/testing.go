package helpers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/thiduzz/slack-bot/models"
	"net/url"
)

func EncodeToBase64URL(request map[string]interface{}) string {

	params := url.Values{}
	for key, value := range request {
		params.Add(key, fmt.Sprintf("%s",value))
	}
	query := params.Encode()
	return base64.StdEncoding.EncodeToString([]byte(query))
}

func GenerateBaseRequest() map[string]interface{} {
	return map[string]interface{}{
		"token":        "O8mkcDKXfmIitPp7RXSX4S1U",
		"team_id":      "T01T72BF15Z",
		"team_domain":  "test",
		"channel_id":   "C01T72BFMFV",
		"channel_name": "general",
		"user_id":      "U01T02LM6DU",
		"user_name":    "test",
		"command":      "hats",
		"text":         "Whatever",
		"api_app_id":   "TestingAppId",
		"response_url": "https://hooks.slack.com/command/blablabla",
	}
}

func GenerateBaseGroup() models.Group {
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

func GenerateBase64Payload(req map[string]string) string {
	values, _ := url.ParseQuery("")
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(req)
	values.Add("payload", reqBodyBytes.String())
	return base64.StdEncoding.EncodeToString([]byte(values.Encode()))
}

func ConvertRequestToByteSlice(request events.APIGatewayProxyRequest) []byte{
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(request)
	return reqBodyBytes.Bytes()
}