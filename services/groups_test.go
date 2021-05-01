package services

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"github.com/thiduzz/slack-bot/helpers"
	"github.com/thiduzz/slack-bot/mocks"
	"github.com/thiduzz/slack-bot/repositories"
	"os"
	"testing"
)


func TestDecodeContentFromBase64(t *testing.T) {
	service := GroupService{}
	r := events.APIGatewayProxyRequest{
		Body: "dG9rZW49Tzhta2NES1hmbUlpdFBwN1JYU1g0UzFVJnRlYW1faWQ9VDAxVDcyQkYxNVomdGVhbV9kb21haW49dGhpYWdvcGVyc29uYS1ydTI4NDM2JmNoYW5uZWxfaWQ9QzAxVDcyQkZNRlYmY2hhbm5lbF9uYW1lPWdlbmVyYWwmdXNlcl9pZD1VMDFUMDJMTTZEVSZ1c2VyX25hbWU9dGhpZHV6ejE0JmNvbW1hbmQ9JTJGc29ydGluZy1oYXQtZ3JvdXAtY3JlYXRlJnRleHQ9JmFwaV9hcHBfaWQ9QTAxVDNQOTRINkgmaXNfZW50ZXJwcmlzZV9pbnN0YWxsPWZhbHNlJnJlc3BvbnNlX3VybD1odHRwcyUzQSUyRiUyRmhvb2tzLnNsYWNrLmNvbSUyRmNvbW1hbmRzJTJGVDAxVDcyQkYxNVolMkYxOTMxMjYxMjU3Njg0JTJGVHZscDVXNkJzNXBLMnhRMUhxalZkM0NHJnRyaWdnZXJfaWQ9MTk0ODg5ODg0MTg0MC4xOTI1MDc5NTExMjAzLmU5YzA2ODdmMTUwOGZmYzljNDI1ZmQwYjNhY2FmNjNj",
	}
	res, err := service.Store(r)
	// assert for not nil (good when you expect something)
	if assert.Nil(t, err) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "Group  for channel general created!", res.Body)

	}
}


func TestValidationErrorWhenGroupNameIsTooShort(t *testing.T) {
	service := GroupService{}
	requestBody := helpers.GenerateBaseRequest()
	requestBody["text"] = "Test"
	body := helpers.EncodeToBase64URL(requestBody)
	r := events.APIGatewayProxyRequest{
		Body: body,
	}
	res, err := service.Store(r)
	// assert for not nil (good when you expect something)
	if assert.Nil(t, err) {

		assert.JSONEq(t, `{"response_type":"ephemeral","text":"Group name should be at least 5 character long"}`, res.Body)

	}
}



func TestInsertIntoDatabaseWhenCreatingGroup(t *testing.T) {
	service := GroupService{}
	requestBody := helpers.GenerateBaseRequest()
	body := helpers.EncodeToBase64URL(requestBody)
	r := events.APIGatewayProxyRequest{
		Body: body,
	}
	res, err := service.Store(r)
	// assert for not nil (good when you expect something)
	if assert.Nil(t, err) {
		assert.JSONEq(t, `{"response_type":"ephemeral","text":"Group name should be at least 5 character long"}`, res.Body)
	}
}


func TestListGroupsFromDatabase(t *testing.T) {

	requestBody := helpers.GenerateBaseRequest()
	body := helpers.EncodeToBase64URL(requestBody)
	groupRepositoryMock := &mocks.GroupRepository{}
	var groups []repositories.GroupListItem

	groups = append(groups, repositories.GroupListItem{
		GroupId:   "123123",
		ChannelId: requestBody["channel_id"],
		Title:     "Test Title",
	})


	groupRepositoryMock.On("IndexByContextReference", fmt.Sprintf("%s:%s", requestBody["team_id"], requestBody["channel_id"])).Return(groups, nil).Once()
	filesystem := os.DirFS("../functions/group")
	service := GroupService{
		GroupRepository: groupRepositoryMock,
		fs: filesystem,
	}
	r := events.APIGatewayProxyRequest{
		Body: body,
	}
	res, err := service.Index(r)
	// assert for not nil (good when you expect something)
	if assert.Nil(t, err) {
		assert.Contains(t, res.Body,  "\"text\": \"Test Title\"")
	}
}


func TestDeleteFromDatabase(t *testing.T) {
	service := GroupService{}
	requestBody := helpers.GenerateBaseRequest()
	body := helpers.EncodeToBase64URL(requestBody)
	r := events.APIGatewayProxyRequest{
		Body: body,
	}
	service.Destroy(r)

}