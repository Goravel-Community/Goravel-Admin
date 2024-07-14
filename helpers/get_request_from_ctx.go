package helpers

import (
	"context"

	"github.com/goravel/framework/contracts/http"
)

func GetRequestFromCtx(ctx context.Context) (request *http.ContextRequest) {
	if ctxRequest := ctx.Value("request"); ctxRequest != nil {
		if req, ok := ctx.Value("request").(*http.ContextRequest); ok {
			request = req
		}
		if req, ok := ctx.Value("request").(http.ContextRequest); ok {
			request = &req
		}
	}
	return request
}
