module github.com/gcp-kit/line-bot-boilerplate-go

go 1.11

require (
	cloud.google.com/go/pubsub v1.6.0
	github.com/gcp-kit/line-bot-boilerplate-go/examples/function v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
	github.com/line/line-bot-sdk-go v7.4.0+incompatible
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/gcp-kit/line-bot-boilerplate-go/examples/function => ./examples/function
