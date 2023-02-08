package xss

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"xmirror.cn/iast/goat/util"
)

func ScriptUnSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		input := util.ExtractInput(c, "input")
		t, _ := template.New("foo").Parse(`{{ . }}`)
		t.Execute(c.Writer, template.HTML(input))
	}
}

func ScriptSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		input := util.ExtractInput(c, "input")
		t, _ := template.New("foo").Parse(`{{ . }}`)
		t.Execute(c.Writer, input)
	}
}
