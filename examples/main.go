package main

import (
	"log"
	"os"

	"github.com/gcp-kit/line-bot-boilerplate-go/cmd"
	"github.com/gcp-kit/line-bot-boilerplate-go/constant"
	"github.com/gcp-kit/line-bot-boilerplate-go/examples/function"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"gopkg.in/yaml.v2"
)

type config struct {
	GinMode            string `yaml:"GIN_MODE"`
	BotName            string `yaml:"BOT_NAME"`
	Mid                string `yaml:"MID"`
	ChannelSecret      string `yaml:"CHANNEL_SECRET"`
	ChannelAccessToken string `yaml:"CHANNEL_ACCESS_TOKEN"`
}

func main() {
	fp, err := os.Open("function/.env.yaml")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	cfg := new(config)
	if err := yaml.NewDecoder(fp).Decode(cfg); err != nil {
		panic(err)
	}

	os.Setenv(constant.EnvKeyMid, cfg.Mid)
	os.Setenv(constant.EnvKeyBotName, cfg.BotName)
	os.Setenv(constant.EnvKeyChannelSecret, cfg.ChannelSecret)
	os.Setenv(constant.EnvKeyChannelAccessToken, cfg.ChannelAccessToken)

	tracer := &cmd.Tracer{
		Function: map[constant.TracerName]func(*cmd.Operation, *linebot.Event) *cmd.TracerResp{
			constant.TracerFollowEvent:     function.FollowEvent,
			constant.TracerUnfollowEvent:   function.UnfollowEvent,
			constant.TracerTextMessage:     function.TextEvent,
			constant.TracerLocationMessage: function.LocationEvent,
		},
		LiffFunc: map[string]func(ctx *gin.Context){
			"liff_c": function.Liff, // Compact
		},
	}

	e := gin.Default()
	if err := tracer.Execute(e); err != nil {
		log.Fatal(err)
	}
}
