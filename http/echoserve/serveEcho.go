package echoserve

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Setup() *echo.Echo {
	echo := echo.New()
	echo.Use(middleware.Logger())
	apiGroup := echo.Group("/echo")
	addSsrf(apiGroup.Group("/ssrf"))
	return echo
}
