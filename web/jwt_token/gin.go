package jwt_token

import (
	"context"
	"net/http"

	"github.com/bigbigliu/go_core/database/redis"
	"github.com/bigbigliu/go_core/pkgs"
	"github.com/gin-gonic/gin"
)

func (h *appJWT) TokenVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := pkgs.ResultInfo{}
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			res.Msg = "token不能为空"
			res.Code = "-1"
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		claims, err := h.ParseToken(token)
		if err != nil {
			res.Msg = err.Error()
			res.Code = "-1"
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		if claims == nil {
			res.Msg = "token解析失败"
			res.Code = "-1"
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		username := claims.Username
		token_key := "appos:" + username + ":accesstoken"

		tokenStr, err := redis.Redisclient.Get(context.Background(), token_key).Result()
		if err != nil {
			res.Msg = "token解析失败"
			res.Code = "-1"
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		if tokenStr == "" {
			res.Msg = "token过期"
			res.Code = "-1"
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		c.Set("username", username)
		c.Next()
	}
}
