package cookie

import "net/http"

func CookieHttponlyUnSafe(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:     "cookie",
		Value:    "CookieHttponlyUnSafe",
		HttpOnly: false,
	}
	w.Header().Set("Set-Cookie", c.String())
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte(`{"state":"ok", "msg":"Cookie中的HttpOnly标识会指示浏览器Cookie只能由服务器访问，而不能由js访问。即使攻击者通过控制了页面(xss)，也无法获取敏感Cookie,加大攻击难度。"}`))
}

func CookieHttponlySafe(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:     "cookie",
		Value:    "CookieHttponlySafe",
		HttpOnly: true,
	}
	w.Header().Set("Set-Cookie", c.String())
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte(`{"state":"ok", "msg":"如果此Cookie不需要前端js来操作，只由后端服务器读取，则建议设置HttpOnly"}`))
}
