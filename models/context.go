package models

const ContextTableName = "SortingHatContexts"

type Context struct {
	ContextId  string  `json:"ContextId" valid:"required~Invalid Workspace/Channel Reference"`
	WorkspaceName string   `json:"WorkspaceName" valid:"minstringlength(5)~Workspace name is invalid"`
	ChannelName  string `json:"ChannelName" valid:"required~This is only possible in channels"`
	ChannelId 	string `json:"ChannelId" valid:"required~This is only possible in channels"`
	Groups      []Group `json:"Groups""`
}