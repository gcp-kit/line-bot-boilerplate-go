package tests

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"testing"

	"github.com/gcp-kit/line-bot-boilerplate-go/util"
)

func TestValidateSignature(t *testing.T) {
	var reqBody = `{
		"events": [
			{
				"replyToken": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
				"type": "message",
				"mode": "active",
				"timestamp": 1462629479859,
				"source": {
					"type": "user",
					"userId": "UXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
				},
				"message": {
					"id": "114514",
					"type": "text",
					"text": "Hello, world"
				}
			}
		]
	}`
	secret := "TestSecret"
	body := []byte(reqBody)
	mac := hmac.New(sha256.New, []byte(secret))
	// nolint
	mac.Write(body)
	sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	ok := util.ValidateSignature(secret, sign, body)
	AssertEquals(t, "", ok, true)

	ng := util.ValidateSignature("InvalidSecret", sign, body)
	AssertEquals(t, "", ng, false)
}
