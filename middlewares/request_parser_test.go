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
		Body:                            helpers.GenerateBase64Payload(map[string]string{"type":"view_submission"}),
		IsBase64Encoded:                 true,
	})

		handlerFunction := MiddlewareFunc(ParseRequest(func(ctx context.Context, request *Request) (events.APIGatewayProxyResponse, error){

			assert.NotNil(t, request)
			assert.IsType(t, map[string]interface{}{}, request.DecodedBody)
			assert.Equal(t, "view_submission", request.DecodedBody["type"])
			return events.APIGatewayProxyResponse{}, nil
		}))
		handlerFunction(nil, req)

}
