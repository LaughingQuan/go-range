package response_info

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"xmirror.cn/iast/goat/util"
)

func IDNumberUnSafe() func(c *gin.Context) {

	return func(c *gin.Context) {
		IDNumber := util.ExtractInput(c, "input")

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"username":  "Tom",
				"id_number": IDNumber,
				"birthday":  time.Now().Format("2006-01-02"),
				"nation":    "Han nationality",
			},
		})
	}
}

func IDNumberSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		IDNumber := util.ExtractInput(c, "input")

		// 使用正则对敏感信息做处理
		if util.RegularDetermineIDNumber(IDNumber) {
			IDNumber = IDNumber[:6] + "************"
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"username":  "Tom",
				"id_number": IDNumber,
				"birthday":  time.Now().Format("2006-01-02"),
				"nation":    "Han nationality",
			},
		})
	}
}
