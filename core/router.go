package core

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

// Operation - structure to be managed collectively
type Operation struct {
	ErrMessage []linebot.SendingMessage
	*Tracer
	*linebot.Client
}

// NewRouter - Routing
func (op *Operation) NewRouter(ctx context.Context, engine *gin.Engine) error {
	var wg sync.WaitGroup
	apiGroup := engine.Group("/api")
	{
		apiGroup.POST("/callback", func(c *gin.Context) {
			events, err := op.ParseRequest(c.Request)
			if err != nil {
				if err == linebot.ErrInvalidSignature {
					log.Print(err)
				}
				c.String(http.StatusBadRequest, "NG")
				return
			}
			for _, event := range events {
				// nolint
				wg.Add(1)
				go func(ev *linebot.Event) {
					defer wg.Done()
					err = op.Switcher(ctx, ev)
					if err != nil {
						log.Println(err)
					}
				}(event)
			}
			wg.Wait()
			c.String(http.StatusOK, "OK")
		})
	}

	if len(op.LiffFunc) > 0 {
		for k, v := range op.LiffFunc {
			engine.GET("/"+k, v)
		}
	}

	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Page not found."})
	})
	return nil
}
