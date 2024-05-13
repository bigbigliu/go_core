package web_middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/bigbigliu/go_core/pkgs"

	"github.com/gin-gonic/gin"
)

// IPCounter 用于记录IP请求计数和时间
type IPCounter struct {
	Count      int       // Count ip计数
	LastAccess time.Time // LastAccess 请求时间
}

var (
	ipCounterMap  = make(map[string]*IPCounter)
	mutexIpFilter sync.Mutex
)

// IPFilterMiddleware 中间件用于过滤IP并限制请求频率
// maxRequestsPerIP ip请求阈值 example:200
// timeWindow 统计时间 example: 1*time.Minute
func IPFilterMiddleware(maxRequestsPerIP int, timeWindow time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := pkgs.GetRemoteIP(c)
		mutexIpFilter.Lock()
		counter, exists := ipCounterMap[ip]
		if !exists {
			counter = &IPCounter{}
			ipCounterMap[ip] = counter
		}
		mutexIpFilter.Unlock()

		counter.Count++
		counter.LastAccess = time.Now()

		// 定时器清理过期ip
		go func() {
			time.Sleep(timeWindow)
			mutexIpFilter.Lock()
			defer mutexIpFilter.Unlock()
			if time.Since(counter.LastAccess) >= timeWindow {
				delete(ipCounterMap, ip)
			}
		}()

		// 检查请求频率
		if counter.Count > maxRequestsPerIP {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"code": "-1",
				"msg":  "Too many requests"})
			c.Abort()
			return
		}

		c.Next()
	}
}
