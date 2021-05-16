
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#technical-details">Technical Details</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project

### Interactions

#### User types "/hats"
A retrieval is performed to get all the groups available with a "manage" button
Additionally a button to create a group is displayed

In the backend, retrieve all groups of this context and utilize the Block Kit to generate a nice message:
*  Click on manage will send the <group-id> to the backend and generate a popup with all the group data (tasks, members,dates...)
*  Click on add group will trigger a popup form to fill the group requirements

#### User select a user and click add members on the manage dialog amd click save
A post is performed with all the data of the form and the group_id, in the database an upsert in the groups and tasks table is performed.

## Usage

Run: `make` to generate the binary and then `serverless deploy` to deploy

Run: `yarn build` in the frontend folder and `serverless client deploy` to update the frontend

To generate mocks I am using mockery codegen CLI:

`mockery --dir ./repositories --name MembershipRepository`

## Roadmap

See the [open issues](https://github.com/thiduzz/sorting-hat-slack-bot/issues) for a list of proposed features (and known issues).

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Thiago Mello - [@thizaom](https://twitter.com/thizaom) - thiago.megermello@gmail.com

Project Link: [https://github.com/thiduzz/sorting-hat-slack-bot](https://github.com/thiduzz/sorting-hat-slack-bot)

## Technical Details

###Database Design

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
  
### Slack Interactivity Request
```json
{
  "type": "view_submission",
  "team": {
    "id": "Example",
    "domain": "test-ru28436"
  },
  "user": {
    "id": "Example",
    "username": "exampleuser",
    "name": "Example User",
    "team_id": "Example"
  },
  "api_app_id": "Example",
  "token": "TestToken",
  "trigger_id": "TestTrigger",
  "view": {
    "id": "Example",
    "team_id": "Example",
    "type": "modal",
    "blocks": [
      {
        "type": "divider",
        "block_id": "DRp"
      },
      {
        "type": "section",
        "block_id": "vnS",
        "text": {
          "type": "mrkdwn",
          "text": "No current groups",
          "verbatim": false
        }
      },
      {
        "type": "divider",
        "block_id": "sXnt"
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
    "private_metadata": "T01T72BF15Z:C01T72BFMFV",
    "callback_id": "CreateGroup",
    "state": {
      "values": {
        "inputGroupCreate": {
          "TextInputCreateGroup": {
            "type": "plain_text_input",
            "value": "sadsad"
          }
        }
      }
    },
    "hash": "Example",
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
    "root_view_id": "Example",
    "app_id": "Example",
    "external_id": "",
    "app_installed_team_id": "Example",
    "bot_id": "Example"
  },
  "response_urls": [],
  "is_enterprise_install": false,
  "enterprise": null
}
```