package response_info

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"xmirror.cn/iast/goat/util"
)

func CellUnSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		cell := util.ExtractInput(c, "input")

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"username": "Tom",
				"cell":     cell,
				"type":     "King card",
				"Prize":    "cup",
				"time":    time.Now().Format("2006-01-02"),
			},
		})
	}
}

func CellSafe() func(c *gin.Context) {
	return func(c *gin.Context) {

		cell := util.ExtractInput(c, "input")
		if util.RegularDetermineCell(cell) {
			cell = cell[:3] + "*****" + cell[8:]
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"username": "Tom",
				"cell":     cell,
				"type":     "King card",
				"Prize":    "cup",
				"time":     time.Now().Format("2006-01-02"),
			},
		})
	}
}
