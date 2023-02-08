package irisserve

import (
	"github.com/kataras/iris/v12/core/router"
	"xmirror.cn/iast/goat/targets/iristargets"
	"xmirror.cn/iast/goat/targets/ssti"
)

func addExec(party router.Party) {
	commandGroup := party.Party("/command/{source}")

	commandGroup.Post("/unsafe", iristargets.CommandUnSafe)
	commandGroup.Post("/safe", iristargets.CommandSafe)
	commandGroup.Get("/unsafe", iristargets.CommandUnSafeGet)
	commandGroup.Get("/safe", iristargets.CommandSafeGet)

}
func addSSTI(party router.Party) {
	ExecuteGroup := party.Party("/execute/{source}")
	ExecuteGroup.Get("/unsafe", ssti.ParseUnSafe())
	ExecuteGroup.Get("/safe", ssti.ParseSafe())
}
