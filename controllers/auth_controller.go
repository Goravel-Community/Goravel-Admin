package controllers

import (
	"github.com/goravel-community/goravel-admin/helpers"
	"github.com/goravel-community/goravel-admin/models"
	"github.com/goravel-community/goravel-admin/views/auth"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type AuthController struct {
	// Dependent services
}

func NewAuthController() *AuthController {
	return &AuthController{
		// Inject services
	}
}

func (r *AuthController) Login(ctx http.Context) http.Response {
	errors := ""
	if ctx.Request().Session().Has("error") {
		errors += "<p>" + ctx.Request().Session().Get("error", "").(string) + "</p>"
	}
	// ctx.Response().Header("Content-Type", "text/html")
	// return ctx.Response().Success().String(form)
	return RenderTempl(ctx, auth.Login())
}
func (r *AuthController) PostLogin(ctx http.Context) http.Response {
	email := ctx.Request().Input("email")
	password := ctx.Request().Input("password")
	validator, err := ctx.Request().Validate(map[string]string{
		"email":    "required|email",
		"password": "required",
	})
	if err != nil {
		ctx.Request().Session().Flash("validationError", err.Error())
		return helpers.RedirectBack(ctx)
	}
	if validator.Fails() {
		ctx.Request().Session().Flash("validationError", "Input validation failed")
		ctx.Request().Session().Flash("validationErrors", validator.Errors().All())
		return helpers.RedirectBack(ctx)
	}

	var admin *models.Admin
	facades.Orm().Query().Where("email = ?", email).First(&admin)
	if admin == nil {
		ctx.Request().Session().Flash("validationError", "Email not found")
		return helpers.RedirectBack(ctx)
	}
	if !facades.Hash().Check(password, admin.Password) {
		ctx.Request().Session().Flash("validationError", "Password is invalid")
		return helpers.RedirectBack(ctx)
	}

	token, err := facades.Auth(ctx).Guard("goravel_admin").Login(admin)
	if err != nil {
		ctx.Request().Session().Flash("validationError", err.Error())
		return helpers.RedirectBack(ctx)
	}
	ctx.Request().Session().Put("goravel_admin_token", token)
	ctx.Request().Session().Flash("message", "Welcome back !")

	prefix := facades.Config().Get("goravel_admin.route", "/admin").(string)
	return ctx.Response().Redirect(http.StatusFound, prefix)
}

func (r *AuthController) Logout(ctx http.Context) http.Response {
	facades.Auth(ctx).Guard("goravel_admin").Logout()
	ctx.Request().Session().Forget("goravel_admin_token")

	routePrefix := facades.Config().Get("goravel_admin.route", "/admin").(string)
	return ctx.Response().Redirect(302, routePrefix+"/user/login")
}
