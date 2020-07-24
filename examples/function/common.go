package function

import (
	"errors"
	"fmt"
	"log"

	"github.com/gcp-kit/line-bot-boilerplate-go/cmd"
	"github.com/gcp-kit/line-bot-boilerplate-go/util"
	"github.com/line/line-bot-sdk-go/linebot"
)

// TextEvent - handle text message events
func TextEvent(op *cmd.Operation, event *linebot.Event) *cmd.TracerResp {
	resp := new(cmd.TracerResp)
	message, ok := event.Message.(*linebot.TextMessage)
	if !ok {
		resp.Error = errors.New("couldn't cast")
		return resp
	}
	switch message.Text {
	case "ping":
		items := &linebot.QuickReplyItems{
			Items: []*linebot.QuickReplyButton{
				{
					Action: linebot.QuickReplyAction(&linebot.MessageAction{
						Label: "ping",
						Text:  "ping",
					}),
				},
			},
		}
		prof, err := op.GetProfile(event.Source.UserID).Do()
		if err != nil {
			resp.Error = err
			return resp
		}
		sender := &linebot.Sender{
			Name:    prof.DisplayName,
			IconURL: prof.PictureURL,
		}
		msg := linebot.NewTextMessage("pong").WithQuickReplies(items).WithSender(sender)
		resp.Stack = append(resp.Stack, msg)
	default:
		msg := linebot.NewTextMessage(message.Text)
		resp.Stack = append(resp.Stack, msg)
	}
	return resp
}

// LocationEvent - handle location message events
func LocationEvent(_ *cmd.Operation, event *linebot.Event) *cmd.TracerResp {
	resp := new(cmd.TracerResp)
	message, ok := event.Message.(*linebot.LocationMessage)
	if !ok {
		resp.Error = errors.New("couldn't cast")
		return resp
	}

	text := fmt.Sprintf("Latitude: %f\nLongitude: %f", message.Latitude, message.Longitude)
	resp.Stack = append(resp.Stack, linebot.NewTextMessage(text))
	return resp
}

// FollowEvent - handle follow events
func FollowEvent(op *cmd.Operation, event *linebot.Event) *cmd.TracerResp {
	resp := new(cmd.TracerResp)
	name := util.GetCallFuncName()
	log.Println("Call:", name)

	uid := event.Source.UserID
	log.Println("UID:", uid)

	prof, err := op.GetProfile(uid).Do()
	if err == nil {
		log.Println("Name:", prof.DisplayName)
		log.Println("Picture:", prof.PictureURL)
		text := fmt.Sprintf("%s, thanks for following!", prof.DisplayName)
		resp.Stack = append(resp.Stack, linebot.NewTextMessage(text))
	} else {
		log.Println("Error:", err.Error())
	}
	return resp
}

// UnfollowEvent - handle unfollow events
func UnfollowEvent(_ *cmd.Operation, event *linebot.Event) *cmd.TracerResp {
	resp := new(cmd.TracerResp)
	name := util.GetCallFuncName()
	log.Println("Call:", name)

	uid := event.Source.UserID
	log.Println("UID:", uid)

	return resp
}
