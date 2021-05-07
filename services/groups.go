package services

import (
	"context"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
	"github.com/thiduzz/slack-bot/helpers"
	"github.com/thiduzz/slack-bot/models"
	"github.com/thiduzz/slack-bot/repositories"
	"time"
)

type GroupService struct {
	repositories.GroupRepository
	slack *SlackService
}

func NewGroupService(db *dynamodb.DynamoDB, slackService *SlackService) *GroupService {
	return &GroupService{
		GroupRepository: repositories.NewGroupRepository(db),
		slack:           slackService,
	}
}

func (g GroupService) Index(ctx context.Context, req *models.SlashRequest) (events.APIGatewayProxyResponse, error) {
	_, err := govalidator.ValidateStruct(req.DecodedBody)
	if err != nil {
		return helpers.NewErrorResponse(err), nil
	}
	referenceId := fmt.Sprintf("%s:%s", req.WorkspaceId, req.ChannelId)
	groups, err := g.GroupRepository.IndexByContextReference(referenceId)
	if err != nil {
		return helpers.NewErrorResponse(err), nil
	}
	err = g.slack.showGroupsListDialog(req.TriggerId, referenceId, groups)
	if err != nil {
		return helpers.NewErrorResponse(err), nil
	}
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func (g GroupService) Store(ctx context.Context, req *models.Request) (events.APIGatewayProxyResponse, error) {
	group := models.Group{
		GroupId:     uuid.NewString(),
		ChannelId:   req.DecodedBody.Team.Id,
		ChannelName: req.DecodedBody.Team.Domain,
		CreatorId:   req.DecodedBody.User.Id,
		Title:       "", //req.DecodedBody.State,
		WorkspaceId: "", //req.DecodedBody.Private,
		CreatedAt:   time.Now().UTC().Format(time.RFC3339),
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
