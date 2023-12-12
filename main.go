// main.go
package main

import (
	"github.com/gin-gonic/gin"

	"core/endpoints"
)

func main() {
	//cfg, err := config.LoadConfig()
	//if err != nil {
	//	panic(err)
	//}
	//ctx := context.ModuleCtx
	//ctx.Init(cfg, nil)

	router := gin.Default()
	// 设置所有注册的端点
	endpoints.SetupRoutes(router, nil)

	router.Run() // 在 0.0.0.0:8080 上监听
}
