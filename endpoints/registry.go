// endpoints/endpoints.go
package endpoints

import (
	mCtx "core/context"
	"github.com/gin-gonic/gin"
)

type Endpoint interface {
	Route() string
	Method() string
	Handler(ctx *mCtx.ModuleContext) gin.HandlerFunc
}

var endpoints []Endpoint

func RegisterEndpoint(ep Endpoint) {
	endpoints = append(endpoints, ep)
}

func SetupRoutes(router *gin.Engine, ctx *mCtx.ModuleContext) {
	for _, ep := range endpoints {
		handler := ep.Handler(ctx)
		switch ep.Method() {
		case "GET":
			router.GET(ep.Route(), handler)
		case "POST":
			router.POST(ep.Route(), handler)
			// 其他方法...
		}
	}
}
