package function

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

// Liff Line Front end Framework(v2.1)
func Liff(ctx *gin.Context) {
	H := gin.H{"Title": "LiffCompactTest"}
	latLon := ctx.Request.Form.Get("latLon")
	if len(latLon) > 0 {
		ua := user_agent.New(ctx.Request.UserAgent())
		un := ua.OSInfo().Name
		url := fmt.Sprintf("https://www.google.com/maps/dir/%s/@%s", latLon, latLon)
		if len(un) > 0 {
			if strings.Contains(un, "iPhone") {
				url = fmt.Sprintf("maps://?daddr=%s&dirflg=w", latLon)
			}
		}
		H["URL"] = url
	}
	jump := ctx.Request.Form.Get("jump")
	if len(jump) > 0 {
		H["URL"] = jump
	}
	err := ctx.Request.Form.Get("error")
	if len(err) > 0 {
		H["Error"] = err
	}
	ctx.HTML(http.StatusOK, "liff.html", H)
}
