package filetargets

import (
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"xmirror.cn/iast/goat/util"
)

func OpenFileInner(input string) error {
	f, err := os.OpenFile(input, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return err
	}
	f.Close()
	return nil
}

func OpenFileUnSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		input := util.ExtractInput(c, "input")

		if err := OpenFileInner(input); err != nil {
			util.GinReturnErr("err file operate", 500, err, c)
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

func OpenFileSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		input := util.ExtractInput(c, "input")
		input = url.QueryEscape(input)
		f, err := os.OpenFile(input, os.O_CREATE|os.O_RDWR, 0600)
		if err != nil {
			util.GinReturnErr("err", 500, err, c)
			return
		}
		f.Close()
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"file_name": input,
			},
		})
	}
}
