package beegoserver

import (
	"github.com/beego/beego/v2/server/web"
	"xmirror.cn/iast/goat/targets/insecurity"
)

func AddUnSec(handlers *web.ControllerRegister) {
	handlers.Post("/beego/unsec/:source/rand/unsafe", insecurity.RandUnSafe())
	handlers.Post("/beego/unsec/:source/rand/safe", insecurity.RandSafe())

	handlers.Post("/beego/unsec/:source/hash/unsafe", insecurity.HashUnSafe())
	handlers.Post("/beego/unsec/:source/hash/safe", insecurity.HashSafe())
}
