package util

import (
	"encoding/json"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

// ParseEvents - extract `[]*linebot.Event`
func ParseEvents(data []byte) []*linebot.Event {
	request := &struct {
		Events []*linebot.Event `json:"events"`
	}{}
	if err := json.Unmarshal(data, request); err != nil {
		log.Fatal(err)
	}
	return request.Events
}
