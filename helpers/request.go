package helpers

import (
	"encoding/base64"
	"net/url"
)

func DecodeRequest(body string) url.Values {
	sDec, _ := base64.StdEncoding.DecodeString(body)
	params, _ := url.ParseQuery(string(sDec))
	return params
}
