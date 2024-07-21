package config

import (
	"context"
	"fmt"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

var DefaultGoravelConfig = map[string]any{
	"route":      "/admin",
	"first_page": "/dashboard",
	"name":       "Goravel Admin",
	"sidebar": func(ctx context.Context, req *http.ContextRequest) (sidebar string) {
		sidebar += `
			<a href="` + adminRoute("/dashboard") + `" class="` + linkActiveClass(req, "/dashboard") + `">
				Dashboard
			</a>
			<a href="` + adminRoute("/products") + `" class="` + linkActiveClass(req, "/products") + `">
				Products
			</a>
		`
		return sidebar
	},
}

func init() {
	config := facades.Config()
	config.Add("goravel_admin", DefaultGoravelConfig)
}

func adminRoute(append string) string {
	return facades.Config().
		GetString("goravel_admin.route", "/admin") + append
}
func linkActiveClass(req *http.ContextRequest, url string) (class string) {
	if req != nil {
		fmt.Println("PATH", (*req).Path())
		if (*req).Path() == adminRoute(url) {
			class += "active"
		}
	}
	return class
}
