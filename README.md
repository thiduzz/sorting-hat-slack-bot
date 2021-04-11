Run: `make` to generate the binary and then `serverless deploy` to deploy

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