package routing

import "github.com/gin-gonic/gin"

var routeRegistry []func(*gin.Engine)

func RegisterRoute(registerFunc func(router *gin.Engine)) {
	routeRegistry = append(routeRegistry, registerFunc)
}

func SetupRoutes(router *gin.Engine) {
	for _, register := range routeRegistry {
		register(router)
		println(register)
	}
}
