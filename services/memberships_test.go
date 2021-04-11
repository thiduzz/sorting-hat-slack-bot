package services

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/mock"
	"github.com/thiduzz/slack-bot/helpers"
	"github.com/thiduzz/slack-bot/mocks"
	"testing"
)

type groupRepository struct{
	mock.Mock
}

// DoSomething is a method on MyMockedObject that implements some interface
// and just records the activity, and returns what the Mock object tells it to.
//
// In the real object, this method would do something useful, but since this
// is a mocked object - we're just going to stub it out.
//
// NOTE: This method is not being tested here, code that uses this object is.
func (m *groupRepository) DoSomething(number int) (bool, error) {

	args := m.Called(number)
	return args.Bool(0), args.Error(1)

}

func TestModalGeneration(t *testing.T) {
	requestBody := helpers.GenerateBaseRequest()
	testGroup := helpers.GenerateBaseGroup()
	groupRepositoryMock := &mocks.GroupRepository{}
	groupRepositoryMock.On("FindByNameAndChannel", requestBody["text"], requestBody["channel_id"]).Return(&testGroup, nil).Once()
	service := MembershipService{
		MembershipRepository: &mocks.MembershipRepository{},
		GroupRepository: groupRepositoryMock,
	}
	body := helpers.EncodeToBase64URL(requestBody)
	r := events.APIGatewayProxyRequest{
		Body: body,
	}
	service.Create(r)

}