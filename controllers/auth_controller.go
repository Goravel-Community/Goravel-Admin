package controllers

import (
	"github.com/goravel-community/goravel-admin/helper"
	"github.com/goravel-community/goravel-admin/models"
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
	if ctx.Request().Session().Has("errors") {
		validationErrors := helper.ValidationErrors(ctx.Request().Session().Get("errors"))
		for fieldname, list := range validationErrors {
			errors += "<p>Validation for field " + fieldname + ": "
			for _, errMessage := range list {
				errors += errMessage + "<br>"
			}
			errors += "</p>"
		}
	}
	form := `
		Goravel Admin Login
		<form method="post" action="/admin/user/login">
			<div>` + errors + `</div>
			<input type="text" name="email" placeholder="Email" value="admin@admin.com">
			<input type="password" name="password" placeholder="Password" value="password">
			<input type="submit" value="Login">
		</form>
	`
	ctx.Response().Header("Content-Type", "text/html")
	return ctx.Response().Success().String(form)
}
func (r *AuthController) PostLogin(ctx http.Context) http.Response {
	email := ctx.Request().Input("email")
	password := ctx.Request().Input("password")
	validator, err := ctx.Request().Validate(map[string]string{
		"email":    "required|email",
		"password": "required",
	})
	if err != nil {
		ctx.Request().Session().Flash("error", err.Error())
		return helper.RedirectBack(ctx)
	}
	if validator.Fails() {
		ctx.Request().Session().Flash("error", "Input validation failed")
		ctx.Request().Session().Flash("errors", validator.Errors().All())
		return helper.RedirectBack(ctx)
	}

	var admin *models.Admin
	facades.Orm().Query().Where("email = ?", email).First(&admin)
	if admin == nil {
		ctx.Request().Session().Flash("error", "Email not found")
		return helper.RedirectBack(ctx)
	}
	if !facades.Hash().Check(password, admin.Password) {
		ctx.Request().Session().Flash("error", "Password is invalid")
		return helper.RedirectBack(ctx)
	}

	token, err := facades.Auth(ctx).Guard("goravel_admin").Login(admin)
	if err != nil {
		ctx.Request().Session().Flash("error", err.Error())
		return helper.RedirectBack(ctx)
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
