package exectargets

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
	"xmirror.cn/iast/goat/util"
)

func CommandUnsafeInner(input string) (string, error) {
	args := shellArgs(input)
	if len(args) == 0 {
		return "", fmt.Errorf("invalid input")
	}

	var cmd *exec.Cmd
	var out bytes.Buffer
	cmd = exec.Command(args[0], args[1:]...)
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

func CommandUnSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
		input := util.ExtractInput(c, "input")
		out, err := CommandUnsafeInner(input)
		if err != nil {
			util.GinReturnErr("Error run command", 500, err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{"out": out},
		})
	}
}

func CommandSafe() func(c *gin.Context) {
	return func(c *gin.Context) {
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
		cmd = exec.Command(input)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			util.GinReturnErr("could not run command", 500, err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{"out": out.String()},
		})

	}
}

func shellArgs(in string) []string {
	return strings.Fields(in)
}
