package goravel_admin

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/goravel-community/goravel-admin/controllers"
	"github.com/goravel-community/goravel-admin/middleware"
	routeFacade "github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
)

//go:embed public/assets/*
var publicAssetsFS embed.FS

// Import the missing package

func RegisterRoutes(route routeFacade.Router, prefix string) {
	dashboard := controllers.NewDashboardController()
	auth := controllers.NewAuthController()
	myAccount := controllers.NewMyAccountController()

	facades.Route().
		Prefix(prefix).
		Middleware(middleware.SendRequestCtx()).
		Group(func(router routeFacade.Router) {
			// serve static assets
			publicAssetsFSPrefix, _ := fs.Sub(publicAssetsFS, "public/assets")
			router.StaticFS("assets", http.FS(publicAssetsFSPrefix))
			// guest routes
			router.Get("/user/login", auth.Login)
			router.Post("/user/login", auth.PostLogin)

			router.Middleware(middleware.Authenticated()).
				Group(func(router routeFacade.Router) {
					// auth routes
					router.Get("/", dashboard.RedirectFirstPage)
					router.Get("/user/logout", auth.Logout)
					// my account routes
					router.Get("/user/profile", myAccount.MyAccountPage)
					router.Post("/user/profile/info", myAccount.SaveInfo)
					router.Post("/user/profile/password", myAccount.SavePassword)
					// dashboard routes
					router.Get("/dashboard", dashboard.Index)
				})
		})

}
