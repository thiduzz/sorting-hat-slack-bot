package middlewares

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"net/url"
)

type ParseRequestHandler func(ctx context.Context, data *Request) (events.APIGatewayProxyResponse, error)

type Request struct {
	events.APIGatewayProxyRequest
	DecodedBody                 map[string]interface{}                          `json:"decodedBody,omitempty"`
}

func ParseRequest(h ParseRequestHandler) HandlerFunc {
	return HandlerFunc(func(ctx context.Context, data json.RawMessage) (interface{}, error) {
		var request Request
		err := json.Unmarshal(data, &request)
		if err != nil{
			return nil, err
		}
		sDec, _ := base64.StdEncoding.DecodeString(request.Body)
		params, _ := url.ParseQuery(string(sDec))
		var body map[string]interface{}
		err = json.Unmarshal([]byte(params.Get("payload")), &body)
		if err != nil{
			return nil, err
		}
		request.DecodedBody = body
		return h(ctx, &request)
	})
}

