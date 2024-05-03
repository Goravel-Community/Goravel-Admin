package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type DashboardController struct {
	// Dependent services
}

func NewDashboardController() *DashboardController {
	return &DashboardController{
		// Inject services
	}
}

func (r *DashboardController) RedirectFirstPage(ctx http.Context) http.Response {
	prefix := facades.Config().Get("goravel_admin.first_page", "/admin").(string)
	first_page := facades.Config().Get("goravel_admin.first_page", "/dashboard").(string)
	return ctx.Response().Redirect(302, prefix+first_page)
}

func (r *DashboardController) Index(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"Hello": "Goravel",
	})
}
