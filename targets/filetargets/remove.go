package filetargets

import (
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"xmirror.cn/iast/goat/util"
)

func RemoveInner(input string) error {
	return os.Remove(input)
}

func RemoveUnSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		input := util.ExtractInput(c, "input")
		if err := RemoveInner(input); err != nil {
			util.GinReturnErr("err", 500, err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"file_name": input,
			},
		})
	}
}

func RemoveSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		input := util.ExtractInput(c, "input")
		input = url.QueryEscape(input)

		f, _ := os.OpenFile(input, os.O_CREATE|os.O_RDWR, 0600)
		f.Close()
		err := os.Remove(input)
		if err != nil {
			util.GinReturnErr("err", 500, err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"file_name": input,
			},
		})
	}
}
