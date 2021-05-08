package services

import (
	"encoding/json"
	"fmt"
	"github.com/slack-go/slack"
	"github.com/thiduzz/slack-bot/models"
	"log"
)

type SlackService struct {
	client *slack.Client
}

func NewSlackService(accessToken string) *SlackService {
	client := slack.New(accessToken)
	return &SlackService{client: client}
}

func (sl SlackService) generateGroupListDialog(referenceId string, groups []models.Group) *slack.ModalViewRequest {

	var groupsBlock []slack.Block
	groupsBlock = append(groupsBlock, slack.NewDividerBlock())
	if len(groups) <= 0 {
		groupsBlock = append(groupsBlock, slack.NewSectionBlock(slack.NewTextBlockObject("plain_text", "No current groups", false, false), nil, nil))
	} else {
		for _, group := range groups {
			groupsBlock = append(groupsBlock,
				slack.NewSectionBlock(
					slack.NewTextBlockObject("plain_text", group.Title, false, false),
					nil,
					&slack.Accessory{
						OverflowElement: slack.NewOverflowBlockElement( "selectOptions|"+ group.GroupId,
							slack.NewOptionBlockObject(
								"manage",
								slack.NewTextBlockObject("plain_text", ":gear: Manage", true, false),
								nil,
							),
							slack.NewOptionBlockObject(
								"delete",
								slack.NewTextBlockObject("plain_text", ":x: Delete", true, false),
								nil,
							),
						),
					}),
			)
		}
	}

	groupsBlock = append(groupsBlock, slack.NewDividerBlock())

	groupsBlock = append(groupsBlock,
		slack.InputBlock{
			Type:           slack.MBTInput,
			BlockID:        "inputGroupCreate",
			Label:          slack.NewTextBlockObject("plain_text", "New group name", false, false),
			Element:        slack.PlainTextInputBlockElement{
				Type:        	slack.METPlainTextInput,
				ActionID:    	"TextInputCreateGroup",
				Placeholder: 	slack.NewTextBlockObject("plain_text", "Write the name....", false, false),
				MinLength: 		5,
				MaxLength: 		25,
			},
			Hint:           nil,
			Optional:       false,
			DispatchAction: true,
		})
	var modalRequest slack.ModalViewRequest
	modalRequest.Type = slack.ViewType("modal")
	modalRequest.Title = slack.NewTextBlockObject("plain_text", "Channel Groups", false, false)
	modalRequest.Close = slack.NewTextBlockObject("plain_text", "Close", false, false)
	modalRequest.Submit = nil
	modalRequest.PrivateMetadata = referenceId
	modalRequest.CallbackID = models.CALLBACK_GROUP_STORE
	modalRequest.Blocks = slack.Blocks{
		BlockSet: groupsBlock,
	}

	return &modalRequest
}

func (sl SlackService) showDialog(triggerId string, modal slack.ModalViewRequest) error {
	body, err := json.Marshal(modal)
	if err != nil {
		log.Println(err)
	}
	log.Println(fmt.Sprintf("View show request: %v", body))

	res, err := sl.client.OpenView(triggerId, modal)
	if err != nil {
		log.Println("error dialog show:"+err.Error())
		log.Println(fmt.Sprintf("View show response: %v", res))
		return err
	}
	log.Println(fmt.Sprintf("View show response: %v", res))
	return nil
}


func (sl SlackService) updateDialog(hashId string, viewId string, modal slack.ModalViewRequest) error {
	body, err := json.Marshal(modal)
	if err != nil {
		log.Println(err)
	}
	log.Println(fmt.Sprintf("View update request: %v", body))

	res, err := sl.client.UpdateView(modal, "", hashId, viewId)
	if err != nil {
		log.Println("error dialog update:"+err.Error())
		log.Println(fmt.Sprintf("View update response: %v", res))
		return err
	}
	log.Println(fmt.Sprintf("View update response: %v", res))
	return nil
}