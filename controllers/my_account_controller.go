package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type MyAccountController struct {
	// Dependent services
}

func NewMyAccountController() *MyAccountController {
	return &MyAccountController{
		// Inject services
	}
}

func (r *MyAccountController) MyAccountPage(ctx http.Context) http.Response {
	return nil
}
func (r *MyAccountController) SaveInfo(ctx http.Context) http.Response {
	return nil
}
func (r *MyAccountController) SavePassword(ctx http.Context) http.Response {
	return nil
}
