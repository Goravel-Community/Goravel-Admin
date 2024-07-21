package controllers

import (
	"github.com/goravel-community/goravel-admin/helpers"
	"github.com/goravel-community/goravel-admin/models"
	"github.com/goravel-community/goravel-admin/views"
	"github.com/goravel/framework/contracts/http"
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
	prefix := helpers.GetConfigString("route")
	first_page := helpers.GetConfigString("first_page")
	return ctx.Response().Redirect(302, prefix+first_page)
}

func (r *DashboardController) Index(ctx http.Context) http.Response {
	welcomeString := "Welcome to the dashboard: guest"
	if admin, ok := ctx.Value("admin").(*models.Admin); ok && admin != nil {
		welcomeString = "Welcome to the dashboard: " + admin.Email
	}
	return RenderTempl(ctx, views.Dashboard(welcomeString))
}

func (r *DashboardController) Products(ctx http.Context) http.Response {
	welcomeString := "hihihihi"
	return RenderTempl(ctx, views.Dashboard(welcomeString))
}
