package irisserve

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func Setup() (*iris.Application, error) {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	routeGroup := app.Party("/iris")
	addExec(routeGroup.Party("/exec"))
	addSSTI(routeGroup.Party("/ssti"))
	err := app.Build()
	return app, err
}
