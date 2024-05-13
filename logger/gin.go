package logger

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GinLogger gin日志请求中间件
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		sugarLogger := Logger

		c.Set("zapLogger", sugarLogger)
		body := ""
		// 检查请求的 Content-Type 是否为 multipart/form-data
		contentType := c.Request.Header.Get("Content-Type")
		if contentType != "multipart/form-data" {
			body, err := io.ReadAll(c.Request.Body)
			if err != nil {
				c.String(http.StatusInternalServerError, "Failed to read request body.")
				return
			}

			// 重新设置请求体，以便后续处理程序可以访问
			c.Request.Body = io.NopCloser(bytes.NewReader(body))
		}

		// 从 Gin 上下文中获取 Request ID
		requestID, exists := c.Get(RequestIDKey)
		var requestIDStr string
		if exists {
			requestIDStr, _ = requestID.(string)
		}

		startTime := time.Now()

		c.Next()

		endTime := time.Now()
		elapsedTime := endTime.Sub(startTime)
		responseTimeMs := float64(elapsedTime.Milliseconds())

		// 请求结束后记录请求信息
		fields := []zap.Field{
			zap.String(RequestIDKey, requestIDStr),
			zap.String("X-Response-Time", fmt.Sprintf("%.2fms", responseTimeMs)),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query-param", c.Request.URL.Query().Encode()),
			zap.Int("status", c.Writer.Status()),
			zap.String("ip", c.ClientIP()),
		}

		// 如果不是表单文件上传，记录 request-body
		if contentType != "multipart/form-data" {
			fields = append(fields, zap.Any("request-body", string(body)))
		}

		// 如果发生错误，记录错误信息
		if len(c.Errors) > 0 {
			errMessages := []string{}
			for _, e := range c.Errors {
				errMessages = append(errMessages, e.Error())
			}
			fields = append(fields, zap.Strings("errors", errMessages))
		}

		// 记录处理程序的信息
		fields = append(fields, zap.String("handler", c.HandlerName()))

		sugarLogger.Info("Request Handled", fields...)
	}
}
