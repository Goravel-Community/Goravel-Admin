package middleware

import (
	"github.com/goravel-community/goravel-admin/models"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func Authenticated() http.Middleware {
	return func(ctx http.Context) {
		var admin *models.Admin
		token, haveToken := ctx.Request().Session().Get("goravel_admin_token", "").(string)
		if haveToken {
			guard := facades.Auth(ctx).Guard("goravel_admin")
			_, err := guard.Parse(token)
			if err == nil {
				_ = guard.User(&admin)
			}
		}
		ctx.WithValue("admin", admin)

		if admin == nil {
			routePrefix := facades.Config().Get("goravel_admin.route", "/admin").(string)
			ctx.Response().Header("Location", routePrefix+"/user/login")
			ctx.Request().AbortWithStatus(http.StatusFound)
			return
		}

		ctx.Request().Next()
	}
}
