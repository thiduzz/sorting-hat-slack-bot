package middlewares

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/thiduzz/slack-bot/models"
	"net/url"
	"strings"
)

type ParseRequestHandler func(ctx context.Context, data *models.InteractivityRequest) (events.APIGatewayProxyResponse, error)

func ParseRequest(h ParseRequestHandler) HandlerFunc {
	return HandlerFunc(func(ctx context.Context, data json.RawMessage) (interface{}, error) {
		var request events.APIGatewayProxyRequest
		err := json.Unmarshal(data, &request)
		if err != nil{
			return nil, err
		}
		sDec, _ := base64.StdEncoding.DecodeString(request.Body)
		params, _ := url.ParseQuery(string(sDec))
		dec := json.NewDecoder(strings.NewReader(params.Get("payload")))
		var decodedBody models.DecodedInteractiveBody
		var proxyRoute *models.ProxyRoute
		if err := dec.Decode(&decodedBody); err != nil {
			return nil, err
		}
		if decodedBody.View.CallbackId != ""{
			callbackSlice := strings.Split(decodedBody.View.CallbackId,".")
			proxyRoute = &models.ProxyRoute{
				Entity: callbackSlice[0],
				Action: callbackSlice[1],
			}
		}else{
			proxyRoute = nil
		}
		return h(ctx, &models.InteractivityRequest{
			APIGatewayProxyRequest: request,
			DecodedInteractiveBody: decodedBody,
			ProxyRoute:             proxyRoute,
		})
	})
}

