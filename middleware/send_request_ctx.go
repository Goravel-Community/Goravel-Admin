package middleware

import (
	"github.com/goravel/framework/contracts/http"
)

func SendRequestCtx() http.Middleware {
	return func(ctx http.Context) {
		ctx.WithValue("request", ctx.Request())
		ctx.Request().Next()
	}
}
