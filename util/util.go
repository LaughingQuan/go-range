package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/v12/context"
	"github.com/labstack/echo/v4"
)

func GetEchoContext(c *gin.Context) echo.Context {
	return echo.New().NewContext(c.Request, c.Writer)
}

func IsFileNotExists(p string) bool {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return true
	}
	return false
}

func ExtractInput(c *gin.Context, key string) string {
	source := c.Param("source")
	switch source {
	case "query":
		return c.Query(key)
	case "buffered-query":
		return buffer(c.Query(key))
	case "params":
		// path parameter includes leading slash, so we chop it off.
		return c.Param(key)[1:]
	case "body":
		return c.PostForm(key)
	case "buffered-body":
		input := c.PostForm(key)
		return buffer(input)
	case "cookies":
		input, err := c.Cookie(key)
		if err != nil {
			c.Error(err)
		}
		return input
	case "headers":
		return c.GetHeader(key)
	case "headers-json":
		// currently only used for SQLi
		inputStr := c.GetHeader(key)
		var creds struct {
			Input string `json:"input"`
		}
		err := json.Unmarshal([]byte(inputStr), &creds)
		if err != nil {
			c.Error(err)
		}
		return creds.Input

	default:
		c.Error(fmt.Errorf("invalid source: %s", source))
		return ""
	}

}

func EchoExtractInput(c echo.Context, key string) string {
	source := c.Param("source")
	switch source {
	case "query":
		return c.QueryParam(key)
	case "buffered-query":
		return buffer(c.QueryParam(key))
	case "params":
		// path parameter includes leading slash, so we chop it off.
		return c.Param(key)[1:]
	case "body":
		return c.FormValue(key)
	case "buffered-body":
		input := c.FormValue(key)
		return buffer(input)
	case "cookies":
		input, err := c.Cookie(key)
		if err != nil {
			c.Error(err)
		}
		return input.Value
	case "headers":
		return c.Request().Header.Get(key)
	case "headers-json":
		// currently only used for SQLi
		inputStr := c.Request().Header.Get(key)
		var creds struct {
			Input string `json:"input"`
		}
		err := json.Unmarshal([]byte(inputStr), &creds)
		if err != nil {
			c.Error(err)
		}
		return creds.Input

	default:
		c.Error(fmt.Errorf("invalid source: %s", source))
		return ""
	}
}

func buffer(s string) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(s)
	return buf.String()
}

func GinReturnErr(msg string, code int, err error, c *gin.Context) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": code,
			"msg":  msg,
			"data": gin.H{
				"err": err.Error(),
			},
		})
	}
}

func EchoReturnErr(msg string, code int, err error, c echo.Context) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": code,
			"msg":  msg,
			"data": gin.H{
				"err": err.Error(),
			},
		})
	}
}
func IrisReturnErr(msg string, code int, err error, c context.Context) {
	if err != nil {
		c.JSON(gin.H{
			"code": code,
			"msg":  msg,
			"data": gin.H{
				"err": err.Error(),
			},
		},
		)
	}
}

func RegularDetermineCell(cellString string) bool {
	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(cellString)
}

func RegularDetermineIDNumber(IDNumber string) bool {
	// (^\d{15}$) 15位身份证
	// (^\d{18}$) 18位身份证
	regRuler := "(^\\d{15}$)|(^\\d{18}$)|(^\\d{17}(\\d|X|x)$)"

	reg := regexp.MustCompile(regRuler)
	return reg.MatchString(IDNumber)
}

func RegularDetermineBankID(BankID string) bool {
	//regRuler := "/^([1-9]{1})(\\d{14}|\\d{18})$/"
	//reg := regexp.MustCompile(regRuler)
	//return reg.MatchString(BankID)
	return true
}

func StartDir() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}
