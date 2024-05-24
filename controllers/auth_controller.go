package controllers

import (
	"github.com/goravel/framework/contracts/http"
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
	return ctx.Response().Success().Json(http.Json{
		"sdsds": "you need to login",
	})
}
