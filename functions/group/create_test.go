package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeContentFromBase64(t *testing.T) {
	r := events.APIGatewayProxyRequest{
		Body: "dG9rZW49Tzhta2NES1hmbUlpdFBwN1JYU1g0UzFVJnRlYW1faWQ9VDAxVDcyQkYxNVomdGVhbV9kb21haW49dGhpYWdvcGVyc29uYS1ydTI4NDM2JmNoYW5uZWxfaWQ9QzAxVDcyQkZNRlYmY2hhbm5lbF9uYW1lPWdlbmVyYWwmdXNlcl9pZD1VMDFUMDJMTTZEVSZ1c2VyX25hbWU9dGhpZHV6ejE0JmNvbW1hbmQ9JTJGc29ydGluZy1oYXQtZ3JvdXAtY3JlYXRlJnRleHQ9JmFwaV9hcHBfaWQ9QTAxVDNQOTRINkgmaXNfZW50ZXJwcmlzZV9pbnN0YWxsPWZhbHNlJnJlc3BvbnNlX3VybD1odHRwcyUzQSUyRiUyRmhvb2tzLnNsYWNrLmNvbSUyRmNvbW1hbmRzJTJGVDAxVDcyQkYxNVolMkYxOTMxMjYxMjU3Njg0JTJGVHZscDVXNkJzNXBLMnhRMUhxalZkM0NHJnRyaWdnZXJfaWQ9MTk0ODg5ODg0MTg0MC4xOTI1MDc5NTExMjAzLmU5YzA2ODdmMTUwOGZmYzljNDI1ZmQwYjNhY2FmNjNj",
	}
	res, err := HandleCreate(r)
	// assert for not nil (good when you expect something)
	if assert.Nil(t, err) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "Group  for channel general created!", res.Body)

	}
}


func TestValidationErrorWhenGroupNameIsTooShort(t *testing.T) {
	r := events.APIGatewayProxyRequest{
		Body: "dG9rZW49Tzhta2NES1hmbUlpdFBwN1JYU1g0UzFVJnRlYW1faWQ9VDAxVDcyQkYxNVomdGVhbV9kb21haW49dGhpYWdvcGVyc29uYS1ydTI4NDM2JmNoYW5uZWxfaWQ9QzAxVDcyQkZNRlYmY2hhbm5lbF9uYW1lPWdlbmVyYWwmdXNlcl9pZD1VMDFUMDJMTTZEVSZ1c2VyX25hbWU9dGhpZHV6ejE0JmNvbW1hbmQ9JTJGc29ydGluZy1oYXQtZ3JvdXAtY3JlYXRlJnRleHQ9ZGRkJmFwaV9hcHBfaWQ9QTAxVDNQOTRINkgmaXNfZW50ZXJwcmlzZV9pbnN0YWxsPWZhbHNlJnJlc3BvbnNlX3VybD1odHRwcyUzQSUyRiUyRmhvb2tzLnNsYWNrLmNvbSUyRmNvbW1hbmRzJTJGVDAxVDcyQkYxNVolMkYxOTI1NjI4NTkwNDY3JTJGek5GREhQUHBQWmNDZlNBbzVXZHFFOHVpJnRyaWdnZXJfaWQ9MTk0OTIxODE2NTQ3Mi4xOTI1MDc5NTExMjAzLjE3ODNjZmYyYTQxNDA0MWEyNGVkNDdjMzNlZGJiODli",
	}
	res, err := HandleCreate(r)
	// assert for not nil (good when you expect something)
	if assert.Nil(t, err) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "{\"response_type\":\"ephemeral\",\"text\":\"Group name should be at least 5 character long\"}", res.Body)

	}
}

