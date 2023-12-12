package endpoints

import (
	mCtx "core/context"
	"core/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	RegisterEndpoint(&SubmoduleEndpoint{})
}

type SubmoduleEndpoint struct {
	// 可以添加更多字段，如服务依赖等
}

func (e *SubmoduleEndpoint) Route() string {
	return "/submodule"
}

func (e *SubmoduleEndpoint) Method() string {
	return "GET"
}

func (e *SubmoduleEndpoint) Handler(ctx *mCtx.ModuleContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := handleSubmoduleBusinessLogic(ctx, c.Request.Context())
		utils.Respond(c, result, err) // 使用统一的响应处理
	}
}

func handleSubmoduleBusinessLogic(ctx *mCtx.ModuleContext, req interface{}) (interface{}, error) {
	result := "res"
	return result, nil // 或者在错误时返回 err
}
