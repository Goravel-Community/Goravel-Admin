package goravel_admin

import (
	"github.com/goravel/framework/contracts/config"
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/contracts/route"
)

var (
	App          foundation.Application
	ConfigFacade config.Config
	RouteFacade  route.Route
)

type ServiceProvider struct{}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	App = app

	// app.BindWith(RouteBinding, func(app foundation.Application, parameters map[string]any) (any, error) {
	// 	return NewRoute(app.MakeConfig(), parameters)
	// })
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	ConfigFacade = app.MakeConfig()
	RouteFacade = app.MakeRoute()
	// LogFacade = app.MakeLog()
	// ValidationFacade = app.MakeValidation()
	// ViewFacade = app.MakeView()
	routePrefix := ConfigFacade.Get("goravel_admin.route", "/admin").(string)
	RegisterRoutes(RouteFacade, routePrefix)

	app.Publishes("github.com/goravel-community/goravel-admin", map[string]string{
		"config/goravel_admin.go": app.ConfigPath("goravel_admin.go"),
	}, "goravel-admin-config")

	app.Publishes("github.com/goravel-community/goravel-admin", map[string]string{
		"migrations": app.DatabasePath("migrations"),
	}, "goravel-admin-migrations")
}
