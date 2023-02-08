package unvalidated

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"xmirror.cn/iast/goat/util"
)

func RedirectSafe() func(c *gin.Context) {

	return func(c *gin.Context) {
		formValue := util.ExtractInput(c, "input")
		sanitizedURL := ""
		switch formValue {
		case "1":
			sanitizedURL = "https://www.trust1.com"
		case "2":
			sanitizedURL = "https://rule.trust2.com"
		default:
			sanitizedURL = "https://www.trust.com"
		}
		c.Redirect(http.StatusTemporaryRedirect, sanitizedURL)
	}
}

func RedirectUnSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		formValue := util.ExtractInput(c, "input")

		c.Redirect(http.StatusTemporaryRedirect, formValue)
	}
}
