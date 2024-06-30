package goravel_admin

import (
	"github.com/goravel-community/goravel-admin/controllers"
	"github.com/goravel-community/goravel-admin/middleware"
	routeFacade "github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
)

// Import the missing package

func RegisterRoutes(route routeFacade.Router, prefix string) {
	dashboard := controllers.NewDashboardController()
	auth := controllers.NewAuthController()

	facades.Route().Prefix(prefix).
		Group(func(router routeFacade.Router) {
			// guest routes
			router.Get("/user/login", auth.Login)
			router.Post("/user/login", auth.PostLogin)

			router.Middleware(middleware.Authenticated()).
				Group(func(router routeFacade.Router) {
					// auth routes
					router.Get("/", dashboard.RedirectFirstPage)
					router.Get("/dashboard", dashboard.Index)
					router.Get("/user/logout", auth.Logout)
				})
		})

}
