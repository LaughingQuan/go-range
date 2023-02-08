package sqltargets

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"xmirror.cn/iast/goat/util"
)

func restoreItems(rows *sql.Rows) ([]Vulnerability, error) {
	ret := make([]Vulnerability, 0)

	for rows.Next() {
		var storedID int
		var storedName string
		var storedSeverity int
		var storedDesc string

		if err := rows.Scan(&storedID, &storedName, &storedSeverity, &storedDesc); err != nil {
			return nil, fmt.Errorf("restoreItems: %s", err)
		}
		ret = append(ret, Vulnerability{
			ID:       storedID,
			Name:     storedName,
			Severity: storedSeverity,
			Desc:     storedDesc,
		})
	}

	return ret, nil
}

func SqlQueryInner(key string, querySafely bool) ([]Vulnerability, error) {
	var rows *sql.Rows
	var err error
	if querySafely {
		sql := "SELECT * FROM vulnerability WHERE id = ?;"
		rows, err = GetDB().Raw(sql, key).Rows()
	} else {
		sql := fmt.Sprintf("SELECT * FROM vulnerability WHERE id = %s;", key)
		rows, err = GetDB().Raw(sql).Rows()
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return restoreItems(rows)
}

func SqlQuerySafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := util.ExtractInput(c, "input")

		items, err := SqlQueryInner(id, true)
		if err != nil {
			util.GinReturnErr("err", 500, err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{"items": items},
		})
	}
}

func SqlQueryUnSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := util.ExtractInput(c, "input")

		items, err := SqlQueryInner(id, false)
		if err != nil {
			util.GinReturnErr("err", 500, err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{"items": items},
		})
	}
}
