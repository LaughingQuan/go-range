package chiserver

import (
	"github.com/go-chi/chi/v5"
	"xmirror.cn/iast/goat/targets/cookie"
	"xmirror.cn/iast/goat/targets/xmltargets"
)

func addXPath(r chi.Router) {
	r.Route("/xpath/{source}", func(r chi.Router) {
		r.Post("/unsafe/", xmltargets.XPathUnSafe)
		r.Post("/safe/", xmltargets.XPathSafe)
	})
}

func addCookie(r chi.Router) {
	r.Route("/httponly/{source}", func(r chi.Router) {
		r.Post("/unsafe/", cookie.CookieHttponlyUnSafe)
		r.Post("/safe/", cookie.CookieHttponlySafe)
	})
	r.Route("/cookie/{source}", func(r chi.Router) {
		r.Post("/unsafe/", cookie.CookieUnSafe)
		r.Post("/safe/", cookie.CookieSafe)
	})
}
