package echoserve

import (
	"github.com/labstack/echo/v4"
	"xmirror.cn/iast/goat/targets/ssrf"
)

func addSsrf(ssrfGroup *echo.Group) {
	requestGroup := ssrfGroup.Group("/request/:source")
	requestGroup.POST("/unsafe/*input", ssrf.EchoRequestUnSafe())
	requestGroup.POST("/safe/*input", ssrf.EchoRequestSafe())
}
