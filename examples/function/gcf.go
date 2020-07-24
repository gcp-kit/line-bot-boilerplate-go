package function

// nolint
import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/gcp-kit/line-bot-boilerplate-go/cmd"
	"github.com/gcp-kit/line-bot-boilerplate-go/constant"
	"github.com/gcp-kit/line-bot-boilerplate-go/util"
	"github.com/line/line-bot-sdk-go/linebot"
)

// example
var (
	count int
)

// nolint
func init() {
	if projectID != "" {
		entryPoint = os.Getenv("FUNCTION_NAME")
		setting("parent-test", "child-test")
	}
}

// setFunction add a function to use in `ChildFunctions`
// editing required
func setFunction() {
	tracer.Function = map[constant.TracerName]func(*cmd.Operation, *linebot.Event) *cmd.TracerResp{
		constant.TracerFollowEvent:     FollowEvent,
		constant.TracerUnfollowEvent:   UnfollowEvent,
		constant.TracerTextMessage:     TextEvent,
		constant.TracerLocationMessage: LocationEvent,
	}
	/*
	** the processing to put in the Global variable is here
	 */
	count++ // example
}

// WebHook CloudFunctions(Trigger: HTTP)
// no edit
// nolint
func WebHook(w http.ResponseWriter, r *http.Request) {
	ctx := util.SetGinContext(w, r)
	cmd.WebHook(ctx, secret, parentTopic)
}

// ParentFunctions CloudFunctions(Trigger: Pub/Sub)
// no edit
// nolint
func Forking(_ context.Context, m *pubsub.Message) error {
	log.Println("EntryPoint:", entryPoint)
	switch entryPoint {
	case RouteParentFunctions:
		return cmd.ParentFunctions(m, tracer, childTopic)
	case RouteChildFunctions:
		return cmd.ChildFunctions(m, op)
	default:
		return fmt.Errorf("invalid function name")
	}
}

// LiffFull CloudFunctions(Trigger: HTTP)
// no edit
// nolint
func LiffFull(w http.ResponseWriter, r *http.Request) {
	ctx := util.SetGinContext(w, r)
	Liff(ctx)
}

// LiffTall CloudFunctions(Trigger: HTTP)
// no edit
// nolint
func LiffTall(w http.ResponseWriter, r *http.Request) {
	ctx := util.SetGinContext(w, r)
	Liff(ctx)
}

// LiffCompact CloudFunctions(Trigger: HTTP)
// no edit
// nolint
func LiffCompact(w http.ResponseWriter, r *http.Request) {
	ctx := util.SetGinContext(w, r)
	Liff(ctx)
}
