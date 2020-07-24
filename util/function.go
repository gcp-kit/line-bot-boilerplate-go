package util

import (
	"runtime"
	"strings"
)

// GetCallFuncName - get the function name of the caller
func GetCallFuncName() string {
	skip := 2
	sp := GetCallFuncRoute(skip)
	return sp[len(sp)-1]
}

// GetCallFuncRoute - get the route to the function name of the caller and return it as a slice
func GetCallFuncRoute(skip int) []string {
	// nolint: dogsled
	pc, _, _, _ := runtime.Caller(skip)
	sp := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	return sp
}
