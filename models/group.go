package models

const GroupsTableName = "SortingHatGroups"

type Group struct {
	GroupId  string  `json:"GroupId"`
	Title string   `json:"Title" valid:"minstringlength(10)~Group name should be at least 5 character long"`
	ChannelName  string `json:"ChannelName" valid:"required~This is only possible in channels"`
	ChannelId 	string `json:"ChannelId" valid:"required~This is only possible in channels"`
	WorkspaceId  string `json:"WorkspaceId" valid:"required~Invalid Workspace"`
	CreatorId  string `json:"CreatorId" valid:"required~Invalid User ID"`
	CreatedAt  string `json:"CreatedAt"`
}