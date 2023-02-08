package insecurity

import (
	"crypto/md5"
	"crypto/sha512"
	"fmt"

	"github.com/beego/beego/v2/server/web/context"
)

func HashUnSafe() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		input := ctx.Input.Query("input")
		h := md5.New()
		h.Write([]byte(input))

		ctx.WriteString(fmt.Sprintf(`{"code:200","hashStr":"%x","hashSize":%d}`, h.Sum(nil), h.Size()))
	}
}

func HashSafe() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		input := ctx.Input.Query("input")
		h := sha512.New()
		h.Write([]byte(input))
		ctx.WriteString(fmt.Sprintf(`{"code:200","hashStr":"%x","hashSize":%d}`, h.Sum(nil), h.Size()))
	}
}
