package response_info

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"xmirror.cn/iast/goat/util"
)

func CustomUnSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		input := util.ExtractInput(c, "input")

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"time":    time.Now().Format("2006-01-02"),
				"custom":  input,
				"explain": `The information entered by the user is returned as is. Whether the user-defined matching expression can effectively match the vulnerability information has been detected`,
			},
		})
	}
}

func CustomSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		input := util.ExtractInput(c, "input")
		inputRuneArr := []rune(input)
		for i := len(inputRuneArr)/2 - len(inputRuneArr)/4; i <= len(inputRuneArr)/2+len(inputRuneArr)/4 && i < len(inputRuneArr); i++ {
			inputRuneArr[i] = '*'
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"time":    time.Now().Format("2006-01-02"),
				"custom":  string(inputRuneArr),
				"explain": `The information entered by the user is returned as is. Whether the user-defined matching expression can effectively match the vulnerability information has been detected`,
			},
		})
	}
}
