package helper

import "github.com/goravel/framework/contracts/http"

func RedirectBack(ctx http.Context) http.Response {
	return ctx.Response().Redirect(http.StatusFound, ctx.Request().Header("Referer", "/"))
}
