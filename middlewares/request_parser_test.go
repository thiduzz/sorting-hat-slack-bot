package middlewares

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"github.com/thiduzz/slack-bot/helpers"
	"testing"
)

func TestDecodingOfRequest(t *testing.T) {
	req := helpers.ConvertRequestToByteSlice(events.APIGatewayProxyRequest{
		Body:                            helpers.GenerateBase64Payload(map[string]interface{}{"type":"view_submission","trigger_id":"testtrigger"}),
		IsBase64Encoded:                 true,
	})

		handlerFunction := MiddlewareFunc(ParseRequest(func(ctx context.Context, request *Request) (events.APIGatewayProxyResponse, error){

			assert.NotNil(t, request)
			assert.IsType(t, DecodedBody{}, request.DecodedBody)
			assert.Equal(t, "view_submission", request.DecodedBody.Type)
			return events.APIGatewayProxyResponse{}, nil
		}))
		handlerFunction(nil, req)

}

func TestDefiningCallbackIdAsProxyRoute(t *testing.T) {
	payload := helpers.GenerateBase64Payload(map[string]interface{}{
		"type":"view_submission",
		"view": map[string]string{"callback_id":"group.create"},
	})

	req := helpers.ConvertRequestToByteSlice(events.APIGatewayProxyRequest{
		Body:                            payload,
		IsBase64Encoded:                 true,
	})

	handlerFunction := MiddlewareFunc(ParseRequest(func(ctx context.Context, request *Request) (events.APIGatewayProxyResponse, error){

		assert.NotNil(t, request)
		assert.IsType(t, DecodedBody{}, request.DecodedBody)
		assert.Equal(t, "view_submission", request.DecodedBody.Type)
		return events.APIGatewayProxyResponse{}, nil
	}))
	handlerFunction(nil, req)

}