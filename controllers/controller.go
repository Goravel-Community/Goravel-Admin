package controllers

import (
	"github.com/a-h/templ"
	"github.com/goravel/framework/contracts/http"
)

func RenderTempl(c http.Context, comp templ.Component) http.Response {
	c.Response().Status(200)
	c.Response().Header("Content-Type", "text/html")
	comp.Render(c, c.Response().Writer())
	return nil
}
