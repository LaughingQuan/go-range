package ssti

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/kataras/iris/v12/context"
)

type User struct {
	ID       int
	Username string
	Password string
	Phone    string
}

var strHtml = `<!DOCTYPE html>
<html>

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>go-websocket-study-chat</title>
    <script src="https://cdn.jsdelivr.net/npm/vue@3.0.5/dist/vue.global.js"></script>
    <style>
        .stick-to-b {
            position: sticky;
            bottom: 0px;
        }
    </style>
</head>

<body>
    <div class="stick-to-b">
    <p>当前用户{{.Username}}</p>
    <p>result %s</p>
    </div>
</body>

</html>`

var user = User{0, "Tom", "45dcc35354c0801b", "17758566989"}

func ParseUnSafe() func(ctx context.Context) {
	return func(ctx context.Context) {
		input := ctx.FormValue("input")
		var text = fmt.Sprintf(strHtml, input)
		tmpl := template.New("hello")
		tmpl, _ = tmpl.Parse(text)
		tmpl.Execute(ctx.ResponseWriter(), user)
	}
}

func ParseSafe() func(ctx context.Context) {
	return func(ctx context.Context) {
		input := ctx.FormValue("input")
		if strings.Contains(input, "}}") || strings.Contains(input, "{{") {
			input = "输入带有不规范字符,请尽量避免使用用户输入渲染模板"
		}
		var text = fmt.Sprintf(strHtml, input)
		tmpl := template.New("hello")
		tmpl, _ = tmpl.Parse(text)
		tmpl.Execute(ctx.ResponseWriter(), user)
	}
}
