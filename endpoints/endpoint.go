package endpoints

import (
	"core/di"
	"core/pkg/interfaces/handlers"
	"core/routing"
	"github.com/gin-gonic/gin"
)

func init() {
	routing.RegisterRoute(func(router *gin.Engine) {
		router.GET("/users", getUserHandler())
	})
}

// 采用闭包来将 di.GetServiceLocator(nil) 做一个时序上的后置
func getUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userService := di.GetServiceLocator(nil).UserService()
		handler := handlers.NewUserHandler(userService)
		handler.GetUser(c)
	}
}
