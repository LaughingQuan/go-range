package ssrf

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"xmirror.cn/iast/goat/util"
)

func GinRequestUnSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		input := util.ExtractInput(c, "input")
		url := "http://example.com" + input
		res, err := http.Get(url)
		if err != nil {
			util.GinReturnErr("parameter is wrong", 500, err, c)
			return
		}
		c.DataFromReader(
			http.StatusOK,
			res.ContentLength,
			"text/html",
			res.Body,
			nil,
		)
	}
}

func GinRequestSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		input := util.ExtractInput(c, "input")
		urlStr := "http://example.com" + url.QueryEscape(input)
		res, err := http.Get(urlStr)
		if err != nil {
			util.GinReturnErr("http.Get is wrong", 500, err, c)
			return
		}
		c.DataFromReader(
			http.StatusOK,
			res.ContentLength,
			"text/html",
			res.Body,
			nil,
		)
	}
}

func EchoRequestUnSafe() func(c echo.Context) error {
	return func(c echo.Context) error {
		input := util.EchoExtractInput(c, "input")
		res, err := http.Get(input)
		if err != nil {
			util.EchoReturnErr("http.Get wrong", 500, err, c)
			return err
		}
		buf, err := ioutil.ReadAll(res.Body)
		if err != nil {
			util.EchoReturnErr("ioutil.ReadAll is wrong", 500, err, c)
			return err
		}
		c.HTML(http.StatusOK, string(buf))
		return nil
	}
}

func EchoRequestSafe() func(c echo.Context) error {
	return func(c echo.Context) error {
		input := util.EchoExtractInput(c, "input")
		if !SSRFHostCheck(input) {
			c.String(http.StatusOK, "illegal url"+input)
			return nil
		}
		res, err := http.Get(input)
		if err != nil {
			util.EchoReturnErr("http.Get is wrong", 500, err, c)
			return err
		}
		buf, err := ioutil.ReadAll(res.Body)
		if err != nil {
			util.EchoReturnErr("ioutil.ReadAll is wrong", 500, err, c)
			return err
		}
		c.HTML(http.StatusOK, string(buf))
		return nil
	}
}

func SSRFHostCheck(urlStr string) bool {
	url, err := url.Parse(urlStr)
	if err != nil {
		return false
	}
	host := strings.ToLower(url.Host)
	hostwhitelist := "www.baidu.com" //白名单
	return host == hostwhitelist
}
