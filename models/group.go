package models

const GroupsTableName = "SortingHatGroups"

type Group struct {
	GroupId     			string 	`json:"group_id" valid:"required,uuid~Group needs an ID"`
	ContextReference     	string 	`json:"context_reference" valid:"required~Context reference is missing"`
	Title       			string 	`json:"title" valid:"minstringlength(5)~Group name should be at least 5 character long"`
	Creator					Creator	`json:"creator" valid:"required~Creator is invalid"`
	StartsAt   				string 	`json:"starts_at" valid:"rfc3339~Starts At date format is invalid"`
	EndsAt   				string 	`json:"ends_at" valid:"rfc3339~Ends At date format is invalid"`
	CreatedAt   			string 	`json:"created_at" valid:"rfc3339~Created At date format is invalid"`
}

type Creator struct {
	Name     				string `json:"name" valid:"required~User name is required"`
	SlackId     			string `json:"slack_id" valid:"required~User is invalid"`
}
