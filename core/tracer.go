package core

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

// Tracer - hold processing function for each event
type Tracer struct {
	Function map[TracerName]func(ctx context.Context, op *Operation, event *linebot.Event) *TracerResp
	LiffFunc map[string]func(ctx *gin.Context)
}

// TracerResp - hold items for outgoing messages
type TracerResp struct {
	Stack []linebot.SendingMessage
	Error error
}

// Execute - create and run instance
func (tracer *Tracer) Execute(ctx context.Context, engine *gin.Engine) error {
	secret, ok := os.LookupEnv(EnvKeyChannelSecret)
	if !ok {
		log.Fatalf("no set env [%s]", EnvKeyChannelSecret)
	}

	token, ok := os.LookupEnv(EnvKeyChannelAccessToken)
	if !ok {
		log.Fatalf("no set env [%s]", EnvKeyChannelAccessToken)
	}

	client, err := linebot.New(secret, token)
	if err != nil {
		return err
	}

	op := &Operation{Client: client, Tracer: tracer}

	if err := op.NewRouter(ctx, engine); err != nil {
		return err
	}

	if err := engine.Run(":1234"); err != nil {
		return err
	}
	return nil
}
