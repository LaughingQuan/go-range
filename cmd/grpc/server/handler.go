package main

import (
	"fmt"
	"log"
	"strings"

	pb "xmirror.cn/iast/goat/rpc"
	exe "xmirror.cn/iast/goat/targets/exectargets"
	fil "xmirror.cn/iast/goat/targets/filetargets"
	sql "xmirror.cn/iast/goat/targets/sqltargets"
)

type SampleHandler = func(string) (*pb.GoatResponse, error)

var (
	registry = make(map[string]SampleHandler)
)

func init() {
	registry["cmd"] = cmdHandler
	registry["gen"] = genHandler
	registry["rem"] = remHandler
	registry["sql"] = sqlHandler
}

func handle(cmd string, input string) (*pb.GoatResponse, error) {
	log.Printf("receiving %s input: %s", cmd, input)

	h, ok := registry[cmd]
	if !ok {
		return &pb.GoatResponse{
			Content: "invalid request command: " + cmd,
		}, nil
	}
	return h(input)
}

func sqlHandler(input string) (*pb.GoatResponse, error) {
	items, err := sql.SqlQueryInner(input, false)
	if err != nil {
		return nil, err
	}

	var sb strings.Builder
	for _, item := range items {
		sb.WriteString(fmt.Sprintf("(%d)\t[level=%d]\t%s\t%s", item.ID, item.Severity, item.Name, item.Desc))
	}

	return &pb.GoatResponse{
		Content: sb.String(),
	}, nil
}

func cmdHandler(input string) (*pb.GoatResponse, error) {
	res, err := exe.CommandUnsafeInner(input)
	if err != nil {
		return nil, err
	}

	return &pb.GoatResponse{
		Content: res,
	}, nil
}

func genHandler(input string) (*pb.GoatResponse, error) {
	if err := fil.OpenFileInner(input); err != nil {
		return nil, err
	}

	return &pb.GoatResponse{
		Content: "file create done",
	}, nil
}

func remHandler(input string) (*pb.GoatResponse, error) {
	if err := fil.RemoveInner(input); err != nil {
		return nil, err
	}

	return &pb.GoatResponse{
		Content: "file delete done ",
	}, nil
}
