package goravel_admin

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/onlinedigital/goravel-admin/controllers"
)

func RegisterRoutes(route route.Router, prefix string) {
	router := route.Prefix(prefix)

	dashboard := controllers.NewDashboardController()
	router.Get("/", dashboard.Login)

	auth := controllers.NewAuthController()
	router.Get("/login", auth.Login)

}
