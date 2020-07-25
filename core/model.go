package core

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

// SendReplyMessage - send ReplyMessage
func (op *Operation) SendReplyMessage(token string, messages []linebot.SendingMessage) error {
	if _, err := op.ReplyMessage(token, messages...).Do(); err != nil {
		return err
	}
	return nil
}

// SendPushMessage - send PushMessage
func (op *Operation) SendPushMessage(uid string, messages []linebot.SendingMessage) error {
	if _, err := op.PushMessage(uid, messages...).Do(); err != nil {
		return err
	}
	return nil
}
