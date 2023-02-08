package beegoserver

import "github.com/beego/beego/v2/server/web"

func Setup() *web.ControllerRegister {
	handlers := web.NewControllerRegister()
	AddUnSec(handlers)
	return handlers
}
