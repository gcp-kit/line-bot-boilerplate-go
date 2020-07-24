package tests

import (
	"testing"

	"github.com/gcp-kit/line-bot-boilerplate-go/util"
)

func TestCommonGetCallFunc(t *testing.T) {
	t.Run("GetCallFuncName", func(t *testing.T) {
		AssertEquals(t, "get the function name of the caller", util.GetCallFuncName(), "func1")
	})
	t.Run("GetCallFuncRoute", func(t *testing.T) {
		skip := 1
		route := util.GetCallFuncRoute(skip)
		AssertEquals(t, "get the function name of the caller", route[len(route)-2], "TestCommonGetCallFunc")
	})
}
