package util

import (
	"encoding/json"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/line/line-bot-sdk-go/linebot"
)

// ParseEvents - extract `[]*linebot.Event`
func ParseEvents(m *pubsub.Message) []*linebot.Event {
	request := &struct {
		Events []*linebot.Event `json:"events"`
	}{}
	if err := json.Unmarshal(m.Data, request); err != nil {
		log.Fatal(err)
	}
	return request.Events
}
