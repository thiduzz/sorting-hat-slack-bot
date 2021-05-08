package middlewares

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/gorilla/schema"
	"github.com/thiduzz/slack-bot/models"
	"net/url"
)

type ParseSlashRequestHandler func(ctx context.Context, data *models.SlashRequest) (events.APIGatewayProxyResponse, error)

func ParseSlashRequest(h ParseSlashRequestHandler) HandlerFunc {
	return HandlerFunc(func(ctx context.Context, data json.RawMessage) (interface{}, error) {
		var request events.APIGatewayProxyRequest
		err := json.Unmarshal(data, &request)
		if err != nil{
			return nil, err
		}
		sDec, _ := base64.StdEncoding.DecodeString(request.Body)
		params, _ := url.ParseQuery(string(sDec))
		var decodedBody models.DecodedSlashBody
		schemaDecoder := schema.NewDecoder()
		schemaDecoder.IgnoreUnknownKeys(true)
		err =  schemaDecoder.Decode(&decodedBody, params)
		if err != nil {
			return nil, err
		}
		return h(ctx, &models.SlashRequest{APIGatewayProxyRequest: request, DecodedSlashBody: decodedBody})
	})
}

