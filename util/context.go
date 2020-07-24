package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetGinContext - create `*gin.Context` from `http.ResponseWriter` and `*http.Request`
func SetGinContext(w http.ResponseWriter, r *http.Request) *gin.Context {
	gc, _ := gin.CreateTestContext(w)
	gc.Request = r
	return gc
}
