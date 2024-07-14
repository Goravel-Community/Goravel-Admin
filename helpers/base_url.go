package helpers

import "github.com/goravel/framework/facades"

func AdminRoute(append string) string {
	return facades.Config().
		GetString("goravel_admin.route", "/admin") + append
}
