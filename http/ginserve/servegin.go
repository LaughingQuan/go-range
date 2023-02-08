package ginserve

import (
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	router := gin.Default()
	apiGroup := router.Group("/gin")

	addSqlTarget(apiGroup.Group("/sql"))
	addUnvalidated(apiGroup.Group("/unvalidated"))
	addSsrf(apiGroup.Group("/ssrf"))
	addGinTarget(apiGroup.Group("/ginframework"))
	addXss(apiGroup.Group("/xss"))
	addExec(apiGroup.Group("/exec"))
	addLdap(apiGroup.Group("/ldap"))
	addFileTarget(apiGroup.Group("/file"))
	addResponseInfo(apiGroup.Group("/response-info"))

	return router
}
