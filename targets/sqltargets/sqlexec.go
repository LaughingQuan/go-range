package sqltargets

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"xmirror.cn/iast/goat/util"
)

func SqlExecSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		input := util.ExtractInput(c, "input")
		query := `SELECT "?" as "?"`
		db, err := GetDB().DB()
		if err != nil {
			util.GinReturnErr("err", 500, err, c)
			return
		}
		res, err := db.Exec(query, input, "test")
		if err != nil {
			util.GinReturnErr("err", 500, err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"sql": query,
				"res": res,
			},
		})
	}
}

func SqlExecUnSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		input := util.ExtractInput(c, "input")
		query := fmt.Sprintf(`SELECT "%s" as "%s"`, input, "test")
		db, err := GetDB().DB()
		if err != nil {
			util.GinReturnErr("err", 500, err, c)
			return
		}
		res, err := db.Exec(query)
		if err != nil {
			util.GinReturnErr("err", 500, err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"sql": query,
				"res": res,
			},
		})
	}
}
