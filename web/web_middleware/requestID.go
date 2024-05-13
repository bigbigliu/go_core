package web_middleware

import (
	"github.com/bigbigliu/go_core/pkgs"
	"github.com/gin-gonic/gin"
)

// RequestIDMiddleware requestID请求中间件
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生成唯一的Request ID
		requestID := pkgs.GenUniversalId()

		// 将Request ID添加到请求头
		c.Header("X-Request-ID", requestID)

		// 将Request ID存储在上下文中，以便在请求处理程序中使用
		c.Set("X-Request-ID", requestID)

		c.Next()
	}
}
