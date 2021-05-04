package services

import (
	"encoding/json"
	"github.com/slack-go/slack"
	"github.com/thiduzz/slack-bot/repositories"
	"log"
)

type SlackService struct {
	client *slack.Client
}

func NewSlackService(accessToken string) *SlackService {
	client := slack.New(accessToken)
	return &SlackService{client: client}
}

func (sl SlackService) showGroupsListDialog(triggerId string, referenceId string, groups []repositories.GroupListItem) error {

	var groupsBlock *slack.SectionBlock
	if len(groups) <= 0{
		groupsBlock = slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", "No current groups", false, false), nil,nil)
	}else{
		groupsBlock = slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", "With some groups", false, false), nil,nil)
	}
	var modalRequest slack.ModalViewRequest
	modalRequest.Type = slack.ViewType("modal")
	modalRequest.Title = slack.NewTextBlockObject("plain_text", "Channel Groups", false, false)
	modalRequest.Close = slack.NewTextBlockObject("plain_text", "Close", false, false)
	modalRequest.Submit = slack.NewTextBlockObject("plain_text", "Submit", false, false)
	modalRequest.PrivateMetadata = referenceId
	modalRequest.CallbackID = "CreateGroup"
	modalRequest.Blocks = slack.Blocks{
		BlockSet: []slack.Block{
			slack.NewDividerBlock(),
			groupsBlock,
			slack.NewDividerBlock(),
			slack.InputBlock{
				Type:           "input",
				BlockID:        "inputGroupCreate",
				Label:          slack.NewTextBlockObject("plain_text", "New group name", false, false),
				Element:        slack.NewPlainTextInputBlockElement(slack.NewTextBlockObject("plain_text", "Write the name....", false, false),"TextInputCreateGroup"),
				Hint:           nil,
				Optional:       false,
				DispatchAction: true,
			},
		},
	}

	return sl.showDialog(triggerId, modalRequest)
}

func (sl SlackService) showDialog(triggerId string, modal slack.ModalViewRequest) error  {
	body, err := json.Marshal(modal)
	if err != nil{
		log.Println(err)
	}
	log.Println(string(body))
	log.Println(triggerId)

	_, err = sl.client.OpenView(triggerId, modal)
	if err != nil {
		return err
	}
	return nil
}