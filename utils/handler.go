package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Respond 处理 HTTP 响应的统一方法
func Respond(c *gin.Context, data interface{}, err error) {
	if err != nil {
		// 发生错误时的处理
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data, "status": "ok"})
}
