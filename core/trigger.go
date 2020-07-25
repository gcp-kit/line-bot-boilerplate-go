package core

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"cloud.google.com/go/pubsub"
	"github.com/gcp-kit/line-bot-boilerplate-go/util"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

// WebHook - CloudFunctions(Trigger: HTTP)
func WebHook(ctx *gin.Context, secret string, topic *pubsub.Topic) {
	defer ctx.Request.Body.Close()

	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, "NG")
		return
	}

	if !util.ValidateSignature(secret, ctx.GetHeader("X-Line-Signature"), body) {
		log.Println("ValidateSignature")
		ctx.String(http.StatusBadRequest, "NG")
		return
	}

	msg := &pubsub.Message{Data: body}
	if _, err := topic.Publish(ctx, msg).Get(ctx); err != nil {
		log.Printf("Could not publish message: %v", err)
		ctx.String(http.StatusInternalServerError, "NG")
		return
	}

	ctx.String(http.StatusOK, "OK")
}

// ParentFunctions - CloudFunctions(Trigger: Pub/Sub)
func ParentFunctions(ctx context.Context, message *pubsub.Message, tracer *Tracer, topic *pubsub.Topic) error {
	var wg sync.WaitGroup
	events := tracer.ParseEvents(message)
	for _, event := range events {
		// nolint
		wg.Add(1)
		go func(ev *linebot.Event) {
			defer wg.Done()
			data, err := ev.MarshalJSON()
			if err != nil {
				return
			}
			msg := &pubsub.Message{Data: data}
			if _, err := topic.Publish(ctx, msg).Get(ctx); err != nil {
				return
			}
		}(event)
	}
	wg.Wait()
	return nil
}

// ChildFunctions - CloudFunctions(Trigger: Pub/Sub)
func ChildFunctions(ctx context.Context, message *pubsub.Message, op *Operation) error {
	event := new(linebot.Event)
	if err := event.UnmarshalJSON(message.Data); err != nil {
		return err
	}

	if err := op.Switcher(ctx, event); err != nil {
		return err
	}
	return nil
}
