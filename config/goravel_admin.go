package config

import (
	"github.com/goravel/framework/facades"
)

var DefaultGoravelConfig = map[string]any{
	"route":      "/admin",
	"first_page": "/dashboard",
	"name":       "Goravel Admin",
}

func init() {
	config := facades.Config()
	config.Add("goravel_admin", DefaultGoravelConfig)
}
