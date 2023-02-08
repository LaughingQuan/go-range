package cookie

import (
	"fmt"
	"net/http"
)

func CookieUnSafe(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	input := r.PostForm.Get("input")
	c := http.Cookie{
		Name:     "cookie",
		Value:    input,
		HttpOnly: true,
	}
	w.Header().Set("Set-Cookie", c.String())
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte(fmt.Sprintf(`{"state": "ok", "msg": "应用将未经处理的外部数据存放到Cookie，会导致应用误用未经验证的数据", "input": "%s"}`, input)))

}

func CookieSafe(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	input := r.PostForm.Get("input")

	c := http.Cookie{
		Name:     "cookie",
		Value:    "value",
		HttpOnly: true,
	}
	w.Header().Set("Set-Cookie", c.String())
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte(fmt.Sprintf(`{"state": "ok", "msg": "数据经过处理之后安全存放，避免误用未经验证的数据,此处未使用输入数据作为cookie的值", "input": "%s"}`, input)))
}
