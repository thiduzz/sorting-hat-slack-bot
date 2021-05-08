package middlewares

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"github.com/thiduzz/slack-bot/models"
	"testing"
)

func TestDecodingOfRequest(t *testing.T) {

		handlerFunction := MiddlewareFunc(ParseRequest(func(ctx context.Context, request *models.InteractivityRequest) (events.APIGatewayProxyResponse, error){
			assert.NotNil(t, request)
			assert.IsType(t, &models.ProxyRoute{}, request.ProxyRoute)
			assert.IsType(t, models.DecodedInteractiveBody{}, request.DecodedInteractiveBody)
			assert.Equal(t, "view_submission", request.Type)
			assert.Equal(t,"ThisIsTheGroupName", request.View.State.Values["inputGroupCreate"]["TextInputCreateGroup"].Value)
			return events.APIGatewayProxyResponse{}, nil
		}))
		handlerFunction(nil, models.NewSlackRequest(map[string]interface{}{
			"type":"view_submission",
			"trigger_id":"testtrigger",
			"view": map[string]interface{}{
				"private_metadata":"test",
				"callback_id":"group.create",
				"state": map[string]interface{}{
					"values": map[string] models.SlackStateValuesWrapper{
						"inputGroupCreate": models.SlackStateValuesWrapper{
							"TextInputCreateGroup": models.SlackStateValue{Value: "ThisIsTheGroupName", Type: "plain_text_input"},
						},
					},
				},
			},
		}))

}