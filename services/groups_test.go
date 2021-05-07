package services

/**

func TestValidationErrorWhenGroupNameIsTooShort(t *testing.T) {
	service := GroupService{}
	requestBody := helpers.GenerateBaseRequest()
	requestBody["text"] = "Test"
	r := middlewares.Request{
		DecodedBody: requestBody,
	}
	res, err := service.Store(nil, &r)
	// assert for not nil (good when you expect something)
	if assert.Nil(t, err) {

		assert.JSONEq(t, `{"response_type":"ephemeral","text":"Group name should be at least 5 character long"}`, res.Body)

	}
}

func TestInsertIntoDatabaseWhenCreatingGroup(t *testing.T) {
	service := GroupService{}
	requestBody := helpers.GenerateBaseRequest()
	r := middlewares.Request{
		DecodedBody: requestBody,
	}
	res, err := service.Store(nil, &r)
	// assert for not nil (good when you expect something)
	if assert.Nil(t, err) {
		assert.JSONEq(t, `{"response_type":"ephemeral","text":"Group name should be at least 5 character long"}`, res.Body)
	}
}

func TestListGroupsFromDatabase(t *testing.T) {
	godotenv.Load("../.env")
	requestBody := helpers.GenerateBaseRequest()
	requestBody["trigger_id"] = "123123123"
	r := middlewares.Request{
		DecodedBody: requestBody,
	}
	groupRepositoryMock := &mocks.GroupRepository{}
	var groups []repositories.GroupListItem

	groups = append(groups, repositories.GroupListItem{
		GroupId:   "123123",
		ChannelId: fmt.Sprintf("%s",requestBody["channel_id"]),
		Title:     "Test Title",
	})

	groupRepositoryMock.On("IndexByContextReference", fmt.Sprintf("%s:%s", requestBody["team_id"], requestBody["channel_id"])).Return(groups, nil).Once()
	slack := NewSlackService(os.Getenv("SLACK_ACCESS_TOKEN"))
	service := GroupService{
		GroupRepository: groupRepositoryMock,
		slack:           slack,
	}
	gatewayRequest := events.APIGatewayProxyRequest{
		Body: r.Body,
	}
	res, err := service.Index(gatewayRequest)
	// assert for not nil (good when you expect something)
	if assert.Nil(t, err) {
		assert.Contains(t, res.Body, "\"text\": \"Test Title\"")
	}
}

func TestDeleteFromDatabase(t *testing.T) {
	service := GroupService{}
	requestBody := helpers.GenerateBaseRequest()
	body := helpers.EncodeToBase64URL(requestBody)
	r := events.APIGatewayProxyRequest{
		Body: body,
	}
	service.Destroy(r)

}
*/
