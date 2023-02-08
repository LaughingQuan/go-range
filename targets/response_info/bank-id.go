package response_info

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"xmirror.cn/iast/goat/util"
)

func BankIDUnSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		var bankID = util.ExtractInput(c, "input")

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"username": "Tom",
				"bank-id":  bankID,
				"remarks":  "Black gold member of China testing bank",
				"time":     time.Now().Format("2006-01-02"),
			},
		})
	}
}

func BankIDSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		var bankID = util.ExtractInput(c, "input")

		// 使用正则对敏感信息做处理
		if util.RegularDetermineBankID(bankID) {
			bankID = bankID[:6] + "**********"
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"username": "Tom",
				"bank-id":  bankID,
				"remarks":  "Black gold member of China testing bank",
				"time":     time.Now().Format("2006-01-02"),
			},
		})
	}
}
