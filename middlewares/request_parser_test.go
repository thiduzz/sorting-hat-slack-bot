package middlewares

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"github.com/thiduzz/slack-bot/models"
	"testing"
)

func TestDecodingOfRequest(t *testing.T) {

		handlerFunction := MiddlewareFunc(ParseRequest(func(ctx context.Context, request *models.Request) (events.APIGatewayProxyResponse, error){
			assert.NotNil(t, request)
			assert.IsType(t, models.DecodedBody{}, request.DecodedBody)
			assert.Equal(t, "view_submission", request.DecodedBody.Type)
			return events.APIGatewayProxyResponse{}, nil
		}))
		handlerFunction(nil, models.NewSlackRequest(map[string]interface{}{"type":"view_submission","trigger_id":"testtrigger"}))

}