package pkgs

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// GetRemoteIP 获取客户端ip
func GetRemoteIP(c *gin.Context) string {
	ip := c.Request.Header.Get("X-Real-IP")
	if ip == "" {
		ip = c.Request.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = c.Request.RemoteAddr
	}

	ips := strings.Split(ip, ":")
	if len(ips) > 0 {
		userIp := ips[0]
		if userIp != "" {
			return userIp
		}
		return c.ClientIP()
	}

	return c.ClientIP()
	//return userIp
}
