package goravel_admin

import (
	"github.com/goravel/framework/contracts/config"
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/contracts/route"

	goaConfig "github.com/goravel-community/goravel-admin/config"
)

var (
	App          foundation.Application
	ConfigFacade config.Config
	RouteFacade  route.Route
)

type ServiceProvider struct{}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	App = app
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	ConfigFacade = app.MakeConfig()
	RouteFacade = app.MakeRoute()
	routePrefix := ConfigFacade.GetString("goravel_admin.route", goaConfig.DefaultGoravelConfig["route"].(string))
	RegisterRoutes(RouteFacade, routePrefix)

	app.Publishes("github.com/goravel-community/goravel-admin", map[string]string{
		"config/goravel_admin.go": app.ConfigPath("goravel_admin.go"),
	}, "goravel-admin-config")

	app.Publishes("github.com/goravel-community/goravel-admin", map[string]string{
		"migrations": app.DatabasePath("migrations"),
	}, "goravel-admin-migrations")
}
