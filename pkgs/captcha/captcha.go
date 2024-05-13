package captcha

// 使用 redis自定义 store，支持分布式部署

import (
	"fmt"
	"time"

	cacheCtl "github.com/bigbigliu/go_core/database/redis"
	"github.com/go-redis/redis/v8"
	"github.com/mojocn/base64Captcha"
)

const (
	CaptchaKeyPre = "captcha"
)

// RedisStore 自定义的 Redis 存储器
type RedisStore struct {
	client     *redis.Client
	Expiration time.Duration
}

// NewRedisStore 创建一个新的 Redis 存储器实例
func NewRedisStore(client *redis.Client, expiration time.Duration) *RedisStore {
	return &RedisStore{
		client:     client,
		Expiration: expiration,
	}
}

// base64Captcha.Store Set 将验证码 ID 和对应的值保存到 Redis 中
func (r *RedisStore) Set(id string, value string) error {
	ctx := r.client.Context()
	key := CaptchaKeyPre + ":" + id
	err := r.client.Set(ctx, key, value, r.Expiration).Err()
	if err != nil {
		fmt.Printf("Error setting captcha in Redis: %v\n", err)
		return err
	}

	return nil
}

// base64Captcha.Store Get 根据验证码 ID 从 Redis 中获取对应的值，并可选择是否在获取后清除该值
func (r *RedisStore) Get(id string, clear bool) string {
	ctx := r.client.Context()
	key := CaptchaKeyPre + ":" + id
	value, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	if clear {
		r.client.Del(ctx, key)
	}
	return value
}

// base64Captcha.Store Verify 验证用户输入的答案是否与 Redis 中的值匹配，并可选择是否在验证后清除该值
func (r *RedisStore) Verify(id, answer string, clear bool) bool {
	ctx := r.client.Context()
	key := CaptchaKeyPre + ":" + id
	value, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return false
	}
	if clear {
		r.client.Del(ctx, key)
	}
	return value == answer
}

// base64Captcha_example 使用示例
func base64Captcha_example() {
	// 配置验证码
	var config = base64Captcha.DriverString{
		Height:          80,                                 // 高
		Width:           240,                                // 宽
		NoiseCount:      80,                                 // 背景噪点数量
		ShowLineOptions: 2,                                  // 干扰线数量
		Length:          6,                                  // 验证码长度
		Source:          "123456abcdefghijklmnopqrstuvwxyz", // 验证码随机文本
	}

	// 创建 Redis 存储器
	store := NewRedisStore(cacheCtl.Redisclient, 1*time.Minute)

	// 生成验证码
	captcha := base64Captcha.NewCaptcha(&config, store)

	// 创建验证码
	id, b64s, answer, _ := captcha.Generate()

	// 打印验证码 ID 和答案
	fmt.Println("Captcha ID:", id)
	fmt.Println("Captcha Answer:", answer)
	fmt.Println("Captcha b64s:", b64s)

	captchaVerify := base64Captcha.NewCaptcha(nil, store)
	isValid := captchaVerify.Verify(id, answer, true)
	if !isValid {
		fmt.Println("验证码验证失败")
	}
}
