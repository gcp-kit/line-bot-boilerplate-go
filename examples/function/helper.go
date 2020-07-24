// Package function no edit file
package function

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/gcp-kit/line-bot-boilerplate-go/cmd"
	"github.com/gcp-kit/line-bot-boilerplate-go/constant"
	"github.com/line/line-bot-sdk-go/linebot"
)

const (
	RouteWebHook         = "WebHook"
	RouteParentFunctions = "ParentFunctions"
	RouteChildFunctions  = "ChildFunctions"
	RouteLiffFull        = "LiffFull"
	RouteLiffTall        = "LiffTall"
	RouteLiffCompact     = "LiffCompact"
)

// no edit
var (
	secret       string
	entryPoint   string
	op           *cmd.Operation
	tracer       *cmd.Tracer
	pubSubClient *pubsub.Client
	parentTopic  *pubsub.Topic
	childTopic   *pubsub.Topic
	projectID    = os.Getenv("GCP_PROJECT")
)

// Probably no edit
func setting(parentTopicName, childTopicName string) {
	ctx := context.Background()
	tracer = cmd.NewTracer(ctx)

	var err error
	pubSubClient, err = pubsub.NewClient(tracer.Ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}

	secret = os.Getenv(constant.EnvKeyChannelSecret)
	switch entryPoint {
	case RouteWebHook:
		parentTopic = pubSubClient.Topic(parentTopicName)
	case RouteParentFunctions:
		childTopic = pubSubClient.Topic(childTopicName)
	case RouteChildFunctions:
		setFunction()

		token := os.Getenv(constant.EnvKeyChannelAccessToken)

		client, err := linebot.New(secret, token)
		if err != nil {
			log.Fatal(err)
		}

		op = &cmd.Operation{
			Client: client,
			Tracer: tracer,
		}
	case RouteLiffFull,
		RouteLiffTall,
		RouteLiffCompact:
		// nop
	default:
		// nop
	}
}
