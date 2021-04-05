package services

import (
	"encoding/base64"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func generateBaseRequest() map[string] string{
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

func encodeToBase64URL(request map[string] string) string {

	params := url.Values{}
	for key, value := range request {
		params.Add(key,value)
	}
	query := params.Encode()
	return base64.StdEncoding.EncodeToString([]byte(query))
}

func TestDecodeContentFromBase64(t *testing.T) {
	service := GroupService{}
	r := events.APIGatewayProxyRequest{
		Body: "dG9rZW49Tzhta2NES1hmbUlpdFBwN1JYU1g0UzFVJnRlYW1faWQ9VDAxVDcyQkYxNVomdGVhbV9kb21haW49dGhpYWdvcGVyc29uYS1ydTI4NDM2JmNoYW5uZWxfaWQ9QzAxVDcyQkZNRlYmY2hhbm5lbF9uYW1lPWdlbmVyYWwmdXNlcl9pZD1VMDFUMDJMTTZEVSZ1c2VyX25hbWU9dGhpZHV6ejE0JmNvbW1hbmQ9JTJGc29ydGluZy1oYXQtZ3JvdXAtY3JlYXRlJnRleHQ9JmFwaV9hcHBfaWQ9QTAxVDNQOTRINkgmaXNfZW50ZXJwcmlzZV9pbnN0YWxsPWZhbHNlJnJlc3BvbnNlX3VybD1odHRwcyUzQSUyRiUyRmhvb2tzLnNsYWNrLmNvbSUyRmNvbW1hbmRzJTJGVDAxVDcyQkYxNVolMkYxOTMxMjYxMjU3Njg0JTJGVHZscDVXNkJzNXBLMnhRMUhxalZkM0NHJnRyaWdnZXJfaWQ9MTk0ODg5ODg0MTg0MC4xOTI1MDc5NTExMjAzLmU5YzA2ODdmMTUwOGZmYzljNDI1ZmQwYjNhY2FmNjNj",
	}
	res, err := service.Create(r)
	// assert for not nil (good when you expect something)
	if assert.Nil(t, err) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "Group  for channel general created!", res.Body)

	}
}


func TestValidationErrorWhenGroupNameIsTooShort(t *testing.T) {
	service := GroupService{}
	requestBody := generateBaseRequest()
	requestBody["text"] = "Test"
	body := encodeToBase64URL(requestBody)
	r := events.APIGatewayProxyRequest{
		Body: body,
	}
	res, err := service.Create(r)
	// assert for not nil (good when you expect something)
	if assert.Nil(t, err) {

		assert.JSONEq(t, `{"response_type":"ephemeral","text":"Group name should be at least 5 character long"}`, res.Body)

	}
}



func TestInsertIntoDatabaseWhenCreatingGroup(t *testing.T) {
	service := GroupService{}
	requestBody := generateBaseRequest()
	body := encodeToBase64URL(requestBody)
	r := events.APIGatewayProxyRequest{
		Body: body,
	}
	res, err := service.Create(r)
	// assert for not nil (good when you expect something)
	if assert.Nil(t, err) {
		assert.JSONEq(t, `{"response_type":"ephemeral","text":"Group name should be at least 5 character long"}`, res.Body)
	}
}
