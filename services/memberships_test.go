package services

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/thiduzz/slack-bot/helpers"
	"github.com/thiduzz/slack-bot/mocks"
	"testing"
)

func TestModalGeneration(t *testing.T) {
	requestBody := helpers.GenerateBaseRequest()
	testGroup := helpers.GenerateBaseGroup()
	groupRepositoryMock := &mocks.GroupRepository{}
	groupRepositoryMock.On("FindByNameAndChannel", requestBody["text"], requestBody["channel_id"]).Return(&testGroup, nil).Once()
	service := MembershipService{
		MembershipRepository: &mocks.MembershipRepository{},
		GroupRepository:      groupRepositoryMock,
	}
	body := helpers.EncodeToBase64URL(requestBody)
	r := events.APIGatewayProxyRequest{
		Body: body,
	}
	service.Create(r)

}
