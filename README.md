Run: `make` to generate the binary and then `sam deploy` to deploy

Run: `yarn build` in the frontend folder and `serverless client deploy` to update the frontend

To generate mocks I am using mockery codegen CLI: 

`mockery --dir ./repositories --name MembershipRepository`

* ~~create group~~
* ~~list group~~
* ~~delete group~~
* create membership
* list memberships
* delete membership  
* create tasks for group
* delete task
* set calendar for group tasks


#Database Design

* Tables
    *  SortingHatWorkspace
    ```json
      {
        "workspace": "workspace",
        "channels": ["channel1","channel2"]
      } 
     ```
    *  SortingHatContexts
    ```json
      {
        "_primary_key": {
            "composition": "workspace1:channel1"
        },
        "context_id": "workspace1:channel1",
        "workspace": "workspace1",
        "channel": "channel1",
        "groups": [
          {          
            "group_id": "<uuid-generated>",
            "title": "Random Tasks"
          }  
        ] 
      } 
    ```
    * SortingHatGroups
    ```json
      {
            "_primary_key": {
              "composition": "group_id",
              "sort": "context-id-reference"
            },          
            "group_id": "<uuid-generated>",
            "context_reference": "<context-id-reference>",
            "title": "Random Tasks",
            "creator" : {
                "name": "Thiago",
                "slack_id": "something"            
            },
            "start_at": "date",
            "ends_at": "date",
            "periodicy": "mondays at 2pm",
            "members": [
              {
                "name": "Thiago",
                "slack_id": "something"
              },
              ... 
            ],
            "tasks": [
              {            
                "name": "Monitor",
                "description": "Blabla"
              },
              {
                "name": "Tackle Technical Bugs",
                "description": "Blabla"
              }
            ],
            "created_at": "date"
      }
     ```
    
    * SortingHatSortedTasks
    ```json
      {
        "iteration": "1",
        "group_id" : "<group-id-reference>",
        "broadcasted_at": "date",
        "set": [
            {
              "member": {},
              "task": {}
            },
            ...
        ] 
      }
    ```
  
# Interactions

#### User types "/sorting-hats"
A retrieval is performed to get all the groups available with a "manage" button
Additionally a button to create a group is displayed

In the backend, retrieve all groups of this context and utilize the Block Kit to generate a nice message:
*  Click on manage will send the <group-id> to the backend and generate a popup with all the group data (tasks, members,dates...)
*  Click on add group will trigger a popup form to fill the group requirements

#### User select a user and click add members on the manage dialog amd click save
A post is performed with all the data of the form and the group_id, in the database an upsert in the groups and tasks table is performed.

```json
{
    "type": "view_submission",
    "team": {
        "id": "T01T72BF15Z",
        "domain": "thiagopersona-ru28436"
    },
    "user": {
        "id": "U01T02LM6DU",
        "username": "thiduzz14",
        "name": "thiduzz14",
        "team_id": "T01T72BF15Z"
    },
    "api_app_id": "A01T3P94H6H",
    "token": "O8mkcDKXfmIitPp7RXSX4S1U",
    "trigger_id": "2012555251431.1925079511203.372c1a4a2a68240c67089ba336255688",
    "view": {
        "id": "V020CGB18FR",
        "team_id": "T01T72BF15Z",
        "type": "modal",
        "blocks": [
            {
                "type": "divider",
                "block_id": "unoS"
            },
            {
                "type": "section",
                "block_id": "CpM",
                "text": {
                    "type": "mrkdwn",
                    "text": "No current groups",
                    "verbatim": false
                }
            },
            {
                "type": "divider",
                "block_id": "VrHZ"
            },
            {
                "type": "input",
                "block_id": "inputGroupCreate",
                "label": {
                    "type": "plain_text",
                    "text": "New group name",
                    "emoji": true
                },
                "optional": false,
                "dispatch_action": true,
                "element": {
                    "type": "plain_text_input",
                    "action_id": "TextInputCreateGroup",
                    "placeholder": {
                        "type": "plain_text",
                        "text": "Write the name....",
                        "emoji": true
                    },
                    "dispatch_action_config": {
                        "trigger_actions_on": [
                            "on_enter_pressed"
                        ]
                    }
                }
            }
        ],
        "private_metadata": "TestingPrivateMetaData",
        "callback_id": "",
        "state": {
            "values": {
                "inputGroupCreate": {
                    "TextInputCreateGroup": {
                        "type": "plain_text_input",
                        "value": "test"
                    }
                }
            }
        },
        "hash": "1620110512.C0vHSBuL",
        "title": {
            "type": "plain_text",
            "text": "Channel Groups",
            "emoji": true
        },
        "clear_on_close": false,
        "notify_on_close": false,
        "close": {
            "type": "plain_text",
            "text": "Close",
            "emoji": true
        },
        "submit": {
            "type": "plain_text",
            "text": "Submit",
            "emoji": true
        },
        "previous_view_id": null,
        "root_view_id": "V020CGB18FR",
        "app_id": "A01T3P94H6H",
        "external_id": "",
        "app_installed_team_id": "T01T72BF15Z",
        "bot_id": "B01TWLT7MA4"
    },
    "response_urls": [],
    "is_enterprise_install": false,
    "enterprise": null
}
```