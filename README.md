Run: `make` to generate the binary and then `serverless deploy` to deploy

Run: `yarn build` in the frontend folder and `serverless client deploy` to update the frontend

Run local lambda by: `serverless invoke local --function groupCreate`

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

