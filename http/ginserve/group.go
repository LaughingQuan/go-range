package ginserve

import (
	"github.com/gin-gonic/gin"

	"xmirror.cn/iast/goat/targets/exectargets"
	"xmirror.cn/iast/goat/targets/filetargets"
	"xmirror.cn/iast/goat/targets/gintargets"
	"xmirror.cn/iast/goat/targets/ldaptargets"
	"xmirror.cn/iast/goat/targets/response_info"
	"xmirror.cn/iast/goat/targets/sqltargets"
	"xmirror.cn/iast/goat/targets/ssrf"
	"xmirror.cn/iast/goat/targets/unvalidated"
	"xmirror.cn/iast/goat/targets/xss"
)

func addSqlTarget(sqlGroup *gin.RouterGroup) {
	queryGroup := sqlGroup.Group("/query/:source")
	queryGroup.POST("/safe/*id", sqltargets.SqlQuerySafe())
	queryGroup.POST("/unsafe/*id", sqltargets.SqlQueryUnSafe())

	execGroup := sqlGroup.Group("/exec/:source")
	execGroup.POST("/safe/*input", sqltargets.SqlExecSafe())
	execGroup.POST("/unsafe/*input", sqltargets.SqlExecUnSafe())
}

func addUnvalidated(unvalidatedGroup *gin.RouterGroup) {
	redirectGroup := unvalidatedGroup.Group("/redirect/:source")
	redirectGroup.GET("/safe/*input", unvalidated.RedirectSafe())
	redirectGroup.GET("/unsafe/*input", unvalidated.RedirectUnSafe())
}

func addSsrf(ssrfGroup *gin.RouterGroup) {
	requestGroup := ssrfGroup.Group("/request/:source")
	requestGroup.POST("/unsafe/*input", ssrf.GinRequestUnSafe())
	requestGroup.POST("/safe/*input", ssrf.GinRequestSafe())
}

func addGinTarget(ginGroup *gin.RouterGroup) {
	ginfileGroup := ginGroup.Group("/ginfile/:source")
	ginfileGroup.POST("/unsafe/*input", gintargets.FileUnSafe())
	ginfileGroup.POST("/safe/*input", gintargets.FileSafe())
}

func addXss(xssGroup *gin.RouterGroup) {
	scriptGroup := xssGroup.Group("/script/:source")
	scriptGroup.GET("/unsafe/*input", xss.ScriptUnSafe())
	scriptGroup.GET("/safe/*input", xss.ScriptSafe())

}

func addExec(execGroup *gin.RouterGroup) {
	commandGroup := execGroup.Group("/command/:source")
	commandGroup.POST("/unsafe/*input", exectargets.CommandUnSafe())
	commandGroup.POST("/safe/*input", exectargets.CommandSafe())

	CommandContextGroup := execGroup.Group("/commandctx/:source")
	CommandContextGroup.POST("/unsafe/*input", exectargets.CommandCtxUnSafe())
	CommandContextGroup.POST("/safe/*input", exectargets.CommandCtxSafe())
}

func addLdap(ldapGroup *gin.RouterGroup) {
	secarhGroup := ldapGroup.Group("/search/:source")
	secarhGroup.POST("/unsafe/*input", ldaptargets.SearchUnSafe())
	secarhGroup.POST("/safe/*input", ldaptargets.SearchSafe())
}

func addResponseInfo(responseInfoGroup *gin.RouterGroup) {
	cellGroup := responseInfoGroup.Group("/cell/:source")
	cellGroup.POST("/unsafe/*input", response_info.CellUnSafe())
	cellGroup.POST("/safe/*input", response_info.CellSafe())

	IDNumberGroup := responseInfoGroup.Group("/id-number/:source")
	IDNumberGroup.POST("/unsafe/*input", response_info.IDNumberUnSafe())
	IDNumberGroup.POST("/safe/*input", response_info.IDNumberSafe())

	BankIDGroup := responseInfoGroup.Group("/bank-id/:source")
	BankIDGroup.POST("/unsafe/*input", response_info.BankIDUnSafe())
	BankIDGroup.POST("/safe/*input", response_info.BankIDSafe())

	customGroup := responseInfoGroup.Group("/custom/:source")
	customGroup.POST("/unsafe/*input", response_info.CustomUnSafe())
	customGroup.POST("/safe/*input", response_info.CustomSafe())
}

func addFileTarget(fileGroup *gin.RouterGroup) {
	mkdirGroup := fileGroup.Group("/mkdir/:source")
	mkdirGroup.POST("/unsafe/*input", filetargets.MkdirUnSafe())
	mkdirGroup.POST("/safe/*input", filetargets.MkdirSafe())

	openFileGroup := fileGroup.Group("/openfile/:source")
	openFileGroup.POST("/unsafe/*input", filetargets.OpenFileUnSafe())
	openFileGroup.POST("/safe/*input", filetargets.OpenFileSafe())

	removeGroup := fileGroup.Group("/remove/:source")
	removeGroup.POST("/unsafe/*input", filetargets.RemoveUnSafe())
	removeGroup.POST("/safe/*input", filetargets.RemoveSafe())

	renameGroup := fileGroup.Group("/rename/:source")
	renameGroup.POST("/unsafe/*input", filetargets.RenameUnSafe())
	renameGroup.POST("/safe/*input", filetargets.RenameSafe())

	downloadGroup := fileGroup.Group("/download/:source")
	downloadGroup.POST("/unsafe/*input", filetargets.DownloadUnSafe())
	downloadGroup.POST("/safe/*input", filetargets.DownloadSafe())
}
