package services

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
	"github.com/thiduzz/slack-bot/helpers"
	"github.com/thiduzz/slack-bot/models"
	"github.com/thiduzz/slack-bot/repositories"
	"io/fs"
	"time"
)

type GroupService struct {
	repositories.GroupRepository
	fs fs.FS
}

func NewGroupService(db *dynamodb.DynamoDB, filesystem fs.FS) *GroupService {
	return &GroupService{
		GroupRepository: repositories.NewGroupRepository(db),
		fs: filesystem,
	}
}

func (g GroupService) Index(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	params := helpers.DecodeRequest(request.Body)
	formRequest := map[string]interface{}{"channelId": params.Get("channel_id"), "workspaceId": params.Get("team_id")}
	_, err := govalidator.ValidateMap(formRequest, map[string]interface{}{"channelId":"required,alphanum", "workspaceId": "required,alphanum"})
	if err != nil {
		return helpers.NewErrorResponse(err), nil
	}
	groups, err := g.GroupRepository.IndexByContextReference(fmt.Sprintf("%s:%s", formRequest["workspaceId"],formRequest["channelId"]))
	if err != nil {
		return helpers.NewErrorResponse(err), nil
	}
	formattedResponse, err := helpers.FormatListBlockResponse(g.fs, groups)
	if err != nil {
		return helpers.NewErrorResponse(err), nil
	}
	return events.APIGatewayProxyResponse{Body: formattedResponse, StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":           "application/json",
		}}, nil
}

func (g GroupService) Store(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	params := helpers.DecodeRequest(request.Body)
	group := models.Group{
		GroupId:		uuid.NewString(),
		ChannelId: 		params.Get("channel_id"),
		ChannelName: 	params.Get("channel_name"),
		CreatorId:      params.Get("user_id"),
		Title:        	params.Get("text"),
		WorkspaceId:	params.Get("team_id"),
		CreatedAt:  	time.Now().UTC().Format(time.RFC3339),
	}
	_, err := govalidator.ValidateStruct(group)
	if err != nil {
		return helpers.NewErrorResponse(err), nil
	}
	if err := g.GroupRepository.Store(group); err != nil {
		return helpers.NewErrorResponse(err), nil
	}
	return events.APIGatewayProxyResponse{Body: fmt.Sprintf("New group %s created!", group.Title), StatusCode: 200}, nil
}

func (g GroupService) Destroy(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	params := helpers.DecodeRequest(request.Body)
	title := params.Get("text")
	channel := params.Get("channel_id")
	group, err := g.GroupRepository.FindByNameAndChannel(title, channel)
	if err != nil {
		return helpers.NewErrorResponse(err), nil
	}
	if err := g.GroupRepository.Destroy(group); err != nil {
		return helpers.NewErrorResponse(err), nil
	}
	return events.APIGatewayProxyResponse{Body: fmt.Sprintf("Group sucessfully %s deleted!", title), StatusCode: 200}, nil
}