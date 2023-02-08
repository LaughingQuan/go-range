package exectargets

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"xmirror.cn/iast/goat/util"
)

func CommandCtxUnSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := context.Background()
		var cmd *exec.Cmd
		input := util.ExtractInput(c, "input")
		args := shellArgs(input)
		if len(args) == 0 {
			util.GinReturnErr("input is nil", 500, nil, c)
			return
		}
		cmd = exec.CommandContext(ctx, args[0], args[1:]...)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			util.GinReturnErr("Could not run command", 500, err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{"out": out.String()},
		})

	}
}

func CommandCtxSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := context.Background()
		var cmd *exec.Cmd
		input := util.ExtractInput(c, "input")
		switch input {
		case "ls":
			input = "ls"
		case "whoami":
			input = "whoami"
		default:
			util.GinReturnErr("could not run command", 500, errors.New("an illegal order"), c)
			return
		}
		cmd = exec.CommandContext(ctx, input)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			util.GinReturnErr("Could not run command", 500, err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{"out": out.String()},
		})

	}
}
