package middlewares

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"github.com/thiduzz/slack-bot/models"
	"testing"
)

func TestDecodingOfSlashRequest(t *testing.T) {

		handlerFunction := MiddlewareFunc(ParseSlashRequest(func(ctx context.Context, request *models.SlashRequest) (events.APIGatewayProxyResponse, error){

			assert.NotNil(t, request)
			assert.IsType(t, models.DecodedSlashBody{}, request.DecodedBody)
			assert.Equal(t, "T01T72BF15Z", request.DecodedBody.WorkspaceId)
			assert.Equal(t, "C01T72BFMFV", request.DecodedBody.ChannelId)
			assert.Equal(t, "123123123", request.DecodedBody.TriggerId)
			return events.APIGatewayProxyResponse{}, nil
		}))
		handlerFunction(nil, models.NewSlackRequest(map[string]interface{}{
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
			"trigger_id": 	"123123123",
		}))
}