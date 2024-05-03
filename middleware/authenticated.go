package middleware

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/onlinedigital/goravel-admin/models"
)

func Authenticated() http.Middleware {
	return func(ctx http.Context) {
		guard := facades.Auth().Guard("goravel_admin")
		var admin models.Admin
		err := guard.User(ctx, &admin)
		if err != nil {
			routePrefix := facades.Config().Get("goravel_admin.route", "/admin").(string)
			ctx.Response().Redirect(302, routePrefix+"/login")
			return
		}
		ctx.Request().Next()
	}
}
