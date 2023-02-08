package filetargets

import (
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"xmirror.cn/iast/goat/util"
)

func MkdirUnSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		input := util.ExtractInput(c, "input")
		err := os.MkdirAll(input, os.ModePerm)
		if err != nil {
			util.GinReturnErr("err", 500, err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"folder_name": input,
			},
		})
	}
}

func MkdirSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		input := util.ExtractInput(c, "input")
		input = url.QueryEscape(input)
		err := os.MkdirAll(input, os.ModePerm)
		if err != nil {
			util.GinReturnErr("err", 500, err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"folder_name": input,
			},
		})
	}
}
