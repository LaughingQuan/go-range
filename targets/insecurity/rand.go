package insecurity

import (
	cryptorand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"

	"github.com/beego/beego/v2/server/web/context"
)

func RandUnSafe() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		input := ctx.Input.Query("input")
		var randInt int
		max, err := strconv.Atoi(input)
		if err != nil || max <= 0 {
			randInt = rand.Int()
		} else {
			randInt = rand.Intn(max)
		}
		ctx.WriteString(fmt.Sprintf(`{"code:200","randInt":%d}`, randInt))
	}
}

func RandSafe() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		input := ctx.Input.Query("input")
		var randInt *big.Int
		max, err := strconv.ParseInt(input, 10, 64)
		if err != nil || max <= 0 {
			randInt, _ = cryptorand.Int(cryptorand.Reader, big.NewInt(100))
		} else {
			randInt, _ = cryptorand.Int(cryptorand.Reader, big.NewInt(max))
		}
		ctx.WriteString(fmt.Sprintf(`{"code:200","randInt":%d}`, randInt))
	}
}
