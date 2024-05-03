package goravel_admin

import (
	routeFacade "github.com/goravel/framework/contracts/route"
	"github.com/onlinedigital/goravel-admin/controllers"
	"github.com/onlinedigital/goravel-admin/middleware"
)

// Import the missing package

func RegisterRoutes(route routeFacade.Router, prefix string) {
	adminRoute := func() routeFacade.Router {
		return route.Prefix(prefix)
	}
	authenticatedRoute := func() routeFacade.Router {
		return adminRoute().Middleware(middleware.Authenticated())
	}

	dashboard := controllers.NewDashboardController()
	authenticatedRoute().Get("/", dashboard.RedirectFirstPage)
	authenticatedRoute().Get("/dashboard", dashboard.Index)

	auth := controllers.NewAuthController()
	adminRoute().Get("/login", auth.Login)

}
