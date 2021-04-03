package models

type Group struct {
	GroupId  string  `json:"group_id"`
	Title string   `json:"title" valid:"minstringlength(10)~Group name should be at least 5 character long"`
	ChannelName  string `json:"channel_name" valid:"required~This is only possible in channels"`
	WorkspaceId  string `json:"workspace_id" valid:"required~Invalid Workspace"`
	CreatorId  string `json:"creator_id" valid:"required~Invalid User ID"`
	CreatedAt  string `json:"created_at"`
}