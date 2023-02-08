package main

import (
	"io"
	"net/http"
	"testing"
)

func Test_httpDo(t *testing.T) {
	type args struct {
		url    string
		method string
		reader io.Reader
		header http.Header
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				url: "http://192.168.172.91:8888//iris/ssti/execute/query/safe?input={{.ID}}用户的用户名是 {{.Username }} 密码是{{.Password}} 电话是{{.Phone}}",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := httpDo(tt.args.url, tt.args.method, tt.args.reader, tt.args.header)
			t.Log(got)
		})
	}
}
