package middlewares

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/gorilla/schema"
	"github.com/thiduzz/slack-bot/models"
	"log"
	"net/url"
)

type ParseSlashRequestHandler func(ctx context.Context, data *models.SlashRequest) (events.APIGatewayProxyResponse, error)

func ParseSlashRequest(h ParseSlashRequestHandler) HandlerFunc {
	return HandlerFunc(func(ctx context.Context, data json.RawMessage) (interface{}, error) {
		log.Println(fmt.Sprintf("Request middleware start"))
		var request events.APIGatewayProxyRequest
		err := json.Unmarshal(data, &request)
		if err != nil{
			log.Println(fmt.Sprintf("Request middleware 1 error: %s", err.Error()))
			return nil, err
		}
		sDec, _ := base64.StdEncoding.DecodeString(request.Body)
		params, _ := url.ParseQuery(string(sDec))
		var decodedBody models.DecodedSlashBody
		err =  schema.NewDecoder().Decode(&decodedBody, params)
		if err != nil {
			log.Println(fmt.Sprintf("Request middleware 1 error: %s", err.Error()))
			return nil, err
		}
		log.Println(fmt.Sprintf("Request middleware: %v",request))
		return h(ctx, &models.SlashRequest{APIGatewayProxyRequest: request, DecodedSlashBody: decodedBody})
	})
}

