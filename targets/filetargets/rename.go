package filetargets

import (
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"xmirror.cn/iast/goat/util"
)

func RenameUnSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		input := util.ExtractInput(c, "input")
		err := os.Rename(input, input+".back")
		if err != nil {
			util.GinReturnErr("err", 500, err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"file_name": input,
				"newname":   input + ".back",
			},
		})
	}
}

func RenameSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		input := util.ExtractInput(c, "input")
		input = url.QueryEscape(input)
		err := os.Rename(input, input+"back")
		if err != nil {
			util.GinReturnErr("err", 500, err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"file_name": input,
				"newname":   input + "bak",
			},
		})
	}
}
