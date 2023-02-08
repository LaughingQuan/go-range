package filetargets

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"xmirror.cn/iast/goat/util"
)

func DownloadUnSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		file, handler, err := c.Request.FormFile("input")
		if err != nil {
			util.GinReturnErr("error getting file parameters", 500, err, c)
			return
		}
		defer file.Close()
		timePath, _ := ioutil.TempDir("", "go-test")
		fileName := timePath + "/" + handler.Filename
		f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			util.GinReturnErr("file opening error", 500, err, c)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"file_name":  handler.Filename,
				"store_path": fileName,
			},
		})
	}
}

func DownloadSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		file, handler, err := c.Request.FormFile("input")
		if err != nil {
			util.GinReturnErr("error getting file parameters", 500, err, c)
			return
		}
		defer file.Close()
		if handler.Size > 5<<20 { //In fact, it should be treated more strictly
			util.GinReturnErr("the file is too large", 500, fmt.Errorf("the file size cannot be larger than 5M"), c)
			return
		}
		timePath, _ := ioutil.TempDir("", "go-test")
		fileName := strconv.FormatInt(time.Now().Unix(), 10)
		f, err := os.OpenFile(timePath+"/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			util.GinReturnErr("file opening error", 500, err, c)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"file_name":  handler.Filename,
				"store_path": timePath + "/" + fileName,
			},
		})
	}
}
