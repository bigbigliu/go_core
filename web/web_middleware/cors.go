package web_middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CorsMiddleware 跨域中间件
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		// 核心处理方式
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE,OPTIONS")                                  //允许post访问
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization,Origin,X-Auth-Token,x-requested-with") //header的类型
		c.Header("Access-Control-Max-Age", "1728000")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("content-type", "application/json") //返回数据格式是json

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}
