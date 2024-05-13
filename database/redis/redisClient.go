package redis

import (
	"context"
	"os"
	"strconv"

	"github.com/bigbigliu/go_core/logger"
	goRedis "github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// Redisclient 全局redis客户端
var Redisclient *goRedis.Client

// InitRedisReq 请求参数
type InitRedisReq struct {
	Addr string `json:"addr"` // Addr redis服务host
	Port int    `json:"port"` // Port redis服务port
	Pwd  string `json:"pwd"`  // Pwd redis服务密码
	Db   int    `json:"db"`   // Db redis服务数据库
}

// InitRedis 初始化redis连接
func (h *InitRedisReq) InitRedis() {
	logger.Logger.Info("Redis", zap.String("conn", "connetcing..."))

	addr := h.Addr + ":" + strconv.Itoa(h.Port)
	logger.Logger.Info("Redis", zap.String("redis_addr", addr))
	Redisclient = goRedis.NewClient(&goRedis.Options{
		Addr:     addr,
		Password: h.Pwd,
		DB:       h.Db,
	})

	_, err := Redisclient.Ping(context.Background()).Result()
	if err != nil {
		logger.Logger.Error("Redis", zap.Error(err))
		os.Exit(-1)
	}
	logger.Logger.Info("Redis", zap.String("conn", "Redis连接成功"))
}
