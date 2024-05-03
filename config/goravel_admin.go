package config

import (
	"github.com/goravel/framework/facades"
)

func init() {
	config := facades.Config()
	config.Add("goravel_admin", map[string]any{
		"route":      "/admin",
		"first_page": "/dashboard",
	})
}
