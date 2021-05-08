package services

import (
	"context"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
	"github.com/slack-go/slack"
	"github.com/thiduzz/slack-bot/helpers"
	"github.com/thiduzz/slack-bot/models"
	"github.com/thiduzz/slack-bot/repositories"
	"log"
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

	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		return helpers.NewErrorResponse(err), nil
	}
	referenceId := fmt.Sprintf("%s:%s", req.WorkspaceId, req.ChannelId)
	dialog, err := g.generateGroupDialog(referenceId)
	if err != nil {
		return helpers.NewErrorResponse(err), nil
	}
	err = g.slack.showDialog(req.TriggerId, *dialog)
	if err != nil {
		return helpers.NewErrorResponse(err), nil
	}
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func (g GroupService) generateGroupDialog(referenceId string) (*slack.ModalViewRequest, error) {

	groups, err := g.GroupRepository.IndexByContextReference(referenceId)
	if err != nil {
		return nil, err
	}
	return g.slack.generateGroupListDialog(referenceId, groups), nil
}

func (g GroupService) Store(ctx context.Context, req *models.InteractivityRequest) (events.APIGatewayProxyResponse, error) {
	if _, ok := req.View.State.Values["inputGroupCreate"]["TextInputCreateGroup"]; !ok {
		return helpers.NewInteractivityErrorResponse(map[string]string{"general": "Title is not present"}), nil
	}
	group := models.Group{
		GroupId:     		uuid.NewString(),
		ContextReference: 	req.View.PrivateMetadata,
		Title:       		req.View.State.Values["inputGroupCreate"]["TextInputCreateGroup"].Value,
		Creator:          	models.Creator{
			Name:    req.User.Username,
			SlackId: req.User.Id,
		},
		StartsAt:         	"",
		EndsAt:           	"",
		CreatedAt:   		time.Now().UTC().Format(time.RFC3339),
	}
	_, err := govalidator.ValidateStruct(group)
	if err != nil {
		return helpers.NewInteractivityErrorResponse(map[string]string{"inputGroupCreate": err.Error()}), nil
	}
	if err := g.GroupRepository.Store(group); err != nil {
		log.Println(fmt.Sprintf("Error Storing: %s", err.Error()))
		return helpers.NewErrorResponse(err), nil
	}
	dialog, err := g.generateGroupDialog(req.View.PrivateMetadata)
	if err != nil {
		return helpers.NewErrorResponse(err), nil
	}
	err = g.slack.updateDialog(req.View.Hash, req.View.Id, *dialog)
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
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
