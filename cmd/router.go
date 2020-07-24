package cmd

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

// Operation - structure to be managed collectively
type Operation struct {
	*Tracer
	*linebot.Client
}

// NewRouter - Routing
func (op *Operation) NewRouter(engine *gin.Engine) error {
	var wg sync.WaitGroup
	apiGroup := engine.Group("/api")
	{
		apiGroup.POST("/callback", func(ctx *gin.Context) {
			events, err := op.ParseRequest(ctx.Request)
			if err != nil {
				if err == linebot.ErrInvalidSignature {
					log.Print(err)
				}
				ctx.String(http.StatusBadRequest, "NG")
				return
			}
			for _, event := range events {
				// nolint
				wg.Add(1)
				go func(ev *linebot.Event) {
					defer wg.Done()
					err = op.Switcher(ev)
					if err != nil {
						log.Println(err)
					}
				}(event)
			}
			wg.Wait()
			ctx.String(http.StatusOK, "OK")
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
