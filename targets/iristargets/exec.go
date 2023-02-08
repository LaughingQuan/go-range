package iristargets

import (
	"bytes"
	"errors"
	"os/exec"

	"github.com/kataras/iris/v12"
	"xmirror.cn/iast/goat/util"
)

type H map[string]interface{}

func CommandUnSafe(ctx iris.Context) {
	input := ctx.FormValue("input")
	var cmd = exec.Command(input)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		util.IrisReturnErr("exec fault", 500, err, ctx)
	} else {
		ctx.JSON(
			H{
				"code": 200,
				"msg":  "success",
				"data": H{
					"items": out.String(),
				},
			},
		)
	}
}

func CommandSafe(ctx iris.Context) {
	input := ctx.FormValue("input")
	switch input {
	case "ls":
		input = "ls"
	case "whoami":
		input = "whoami"
	default:
		util.IrisReturnErr("could not run command", 500, errors.New("an illegal order"), ctx)
		return
	}
	var cmd = exec.Command(input)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		util.IrisReturnErr("exec fault", 500, err, ctx)
	} else {
		ctx.JSON(
			H{
				"code": 200,
				"msg":  "success",
				"data": H{
					"items": out.String(),
				},
			},
		)
	}
}

func CommandUnSafeGet(ctx iris.Context) {
	input := &struct {
		Input string `query:input`
	}{}
	ctx.ReadQuery(input)
	var cmd = exec.Command(input.Input)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		util.IrisReturnErr("exec fault", 500, err, ctx)
	} else {
		ctx.JSON(
			H{
				"code": 200,
				"msg":  "success",
				"data": H{
					"items": out.String(),
				},
			},
		)
	}
}

func CommandSafeGet(ctx iris.Context) {
	inputStruct := &struct {
		Input string `query:input`
	}{}
	ctx.ReadQuery(inputStruct)
	var input string
	switch inputStruct.Input {
	case "ls":
		input = "ls"
	case "whoami":
		input = "whoami"
	default:
		util.IrisReturnErr("could not run command", 500, errors.New("an illegal order"), ctx)
		return
	}
	var cmd = exec.Command(input)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		util.IrisReturnErr("exec fault", 500, err, ctx)
	} else {
		ctx.JSON(
			H{
				"code": 200,
				"msg":  "success",
				"data": H{
					"items": out.String(),
				},
			},
		)
	}
}
