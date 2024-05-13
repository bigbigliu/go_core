package web_middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/bigbigliu/go_core/pkgs"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

const redisKeyPrefix = "ip_requests:"

// IPCounterWithRedis 用于记录IP请求计数和时间
type IPCounterWithRedis struct {
	Count      int       // Count ip计数
	LastAccess time.Time // LastAccess 请求时间
}

// IPFilterWithRedisMiddleware 中间件用于过滤IP并限制请求频率(redis)
func IPFilterWithRedisMiddleware(redisClient *redis.Client, maxRequestsPerIP int, timeWindow time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := pkgs.GetRemoteIP(c)

		// 加锁以保护计数器
		mutex := &sync.Mutex{}
		mutex.Lock()
		defer mutex.Unlock()

		countKey := redisKeyPrefix + ip
		lastAccessKey := redisKeyPrefix + ip + "_last_access"

		// 从Redis中获取IP请求计数
		count, _ := redisClient.Get(c, countKey).Int()

		count++
		// 设置IP请求计数，并设置过期时间
		redisClient.Set(c, countKey, count, timeWindow)

		// 获取最后访问时间
		lastAccess, _ := redisClient.Get(c, lastAccessKey).Time()

		// 检查时间窗口内的请求次数
		if count > maxRequestsPerIP && time.Since(lastAccess) < timeWindow {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"code": "-1",
				"msg":  "Too many requests",
			})
			c.Abort()
			return
		}

		// 设置最后访问时间
		redisClient.Set(c, lastAccessKey, time.Now(), timeWindow)

		c.Next()
	}
}
