package cmd

import (
	"context"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

// Switcher - separate processing for each event
func (op *Operation) Switcher(ctx context.Context, event *linebot.Event) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			log.Println(rec)
		}
		switch r := err.(type) {
		case *linebot.APIError:
			if r.Code == http.StatusBadRequest {
				err = nil
			}
		}
	}()

	var eventType TracerName
	switch event.Type {
	case linebot.EventTypeMessage:
		switch message := event.Message.(type) {
		case *linebot.TextMessage,
			*linebot.ImageMessage,
			*linebot.VideoMessage,
			*linebot.AudioMessage,
			*linebot.FileMessage,
			*linebot.LocationMessage,
			*linebot.StickerMessage:
			eventType = reflect.TypeOf(message).Elem().Name()
		default:
			log.Println("invalid MessageType")
			return
		}
	default:
		eventType = string(event.Type)
	}

	fn, ok := op.Function[eventType]
	if !ok {
		log.Printf("no set function [%s]\n", eventType)
		return
	}

	resp := fn(ctx, op, event)
	if resp.Error != nil {
		err = resp.Error
	} else if len(resp.Stack) > 0 {
		err = op.SendReplyMessage(event.ReplyToken, resp.Stack)
	}

	go func() {
		// デバッグモード時にはログを出す
		if gin.IsDebugging() {
			uid := event.Source.UserID
			log.Println("EventType:", eventType)
			log.Println("UID:", uid)
			if prof, _ := op.GetProfile(uid).Do(); prof != nil {
				log.Println("Name:", prof.DisplayName)
				log.Println("Picture:", prof.PictureURL)
			}
		}
	}()

	return err
}
