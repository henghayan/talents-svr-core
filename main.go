// main.go
package main

import (
	"github.com/gin-gonic/gin"

	"core/config"
	"core/di"
	_ "core/endpoints"
	"core/routing"
)

func main() {
	cfg := config.LoadConfig()
	di.GetServiceLocator(cfg)

	router := gin.Default()
	routing.SetupRoutes(router)

	router.Run(cfg.GetAddr())
}
