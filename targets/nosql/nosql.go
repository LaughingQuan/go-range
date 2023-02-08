package nosql

import (
	"context"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"xmirror.cn/iast/goat/util"
)

var once sync.Once

func FindUnSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		once.Do(initData)
		input := util.ExtractInput(c, "input")
		opts := options.Find()
		cursor, err := collection.Find(
			context.TODO(),
			bson.D{{Key: "$where", Value: "this.name == \"" + input + "\""}},
			opts,
		)
		//Bob" || "1"=="1
		if err != nil {
			util.GinReturnErr("Could not query Mongo", 500, err, c)
			return
		}
		var results []bson.M
		if err = cursor.All(context.TODO(), &results); err != nil {
			util.GinReturnErr("Could not query Mongo", 500, err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"items": results,
			},
		})
	}
}

func FindSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		once.Do(initData)
		input := util.ExtractInput(c, "input")
		opts := options.Find()
		cursor, err := collection.Find(
			context.TODO(),
			bson.D{{Key: "name", Value: input}},
			opts,
		)

		if err != nil {
			util.GinReturnErr("Could not query Mongo", 500, err, c)
			return
		}
		var results []bson.M
		if err = cursor.All(context.TODO(), &results); err != nil {
			util.GinReturnErr("Could not query Mongo", 500, err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"items": results,
			},
		})
	}
}
