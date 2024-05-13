package web_middleware

import (
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// TimeoutMiddleware 接口超时中间件
func TimeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(5*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(testResponse),
	)
}

func testResponse(c *gin.Context) {
	c.JSON(http.StatusGatewayTimeout, gin.H{
		"code": http.StatusGatewayTimeout,
		"msg":  "timeout",
	})
}
