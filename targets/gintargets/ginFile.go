package gintargets

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"xmirror.cn/iast/goat/util"
)

func FileSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		input := util.ExtractInput(c, "input")
		if input == "" {
			util.GinReturnErr("err input is nil", 500, fmt.Errorf("input is nil"), c)
			return
		}
		input = url.QueryEscape(input)
		if ok, err := pathExists(input); !ok {
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "success",
				"data": gin.H{
					"err": err,
				},
			})
			return
		}
		c.File(input)
	}
}

func FileUnSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		input := util.ExtractInput(c, "input")
		if input == "" {
			util.GinReturnErr("err input is nil", 500, fmt.Errorf("input is nil"), c)
			return
		}
		if ok, err := pathExists(input); !ok {
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "success",
				"data": gin.H{
					"err": err,
				},
			})
			return
		}
		c.File(input)
	}
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err

}
