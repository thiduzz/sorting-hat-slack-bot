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

type ParseSlashRequestHandler func(ctx context.Context, data *models.SlashRequest) (events.APIGatewayProxyResponse, error)

func ParseSlashRequest(h ParseSlashRequestHandler) HandlerFunc {
	return HandlerFunc(func(ctx context.Context, data json.RawMessage) (interface{}, error) {
		var request models.SlashRequest
		err := json.Unmarshal(data, &request)
		if err != nil{
			return nil, err
		}
		sDec, _ := base64.StdEncoding.DecodeString(request.Body)
		params, _ := url.ParseQuery(string(sDec))
		dec := json.NewDecoder(strings.NewReader(params.Get("payload")))
		var decodedBody models.DecodedSlashBody
		if err := dec.Decode(&decodedBody); err != nil {
			return nil, err
		}
		request.DecodedBody = decodedBody
		request.ChannelId = decodedBody.ChannelId
		request.WorkspaceId = decodedBody.WorkspaceId
		request.TriggerId = decodedBody.TriggerId
		return h(ctx, &request)
	})
}

