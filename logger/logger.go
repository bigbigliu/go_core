package logger

import (
	"context"
	"os"
	"strings"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	RequestIDKey = "X-Request-ID" // RequestIDKey
)

var (
	Logger *zap.Logger
)

// appLog applog
type AppLog struct {
	LogDir        string `json:"log_dir"`         // LogDir 日志保存路径
	LogLevel      string `json:"log_level"`       // LogLevel 日志级别
	ConsoleOutPut bool   `json:"console_out_put"` // ConsoleOutPut 是否输出到控制台
}

// InitializeLogger 初始化日志记录器
func InitializeLogger(param *AppLog) {
	logLevel := zapcore.InfoLevel
	switch strings.ToLower(param.LogLevel) {
	case "debug":
		logLevel = zapcore.DebugLevel
	case "info":
		logLevel = zapcore.InfoLevel
	case "warn":
		logLevel = zapcore.WarnLevel
	case "error":
		logLevel = zapcore.ErrorLevel
	}

	// 配置日志编码器: 定义日志的输出格式以及在日志中显示的字段
	encoderConfig := zapcore.EncoderConfig{
		CallerKey:      "caller_line",                    // 打印文件名和行数
		TimeKey:        "timestamp",                      // 指定时间戳key
		LevelKey:       "level",                          // 指定日志级别key
		NameKey:        "logger",                         //
		MessageKey:     "message",                        // 日志消息文本key
		StacktraceKey:  "stacktrace",                     // 指定日志中堆栈跟踪信息的键名
		LineEnding:     zapcore.DefaultLineEnding,        //
		EncodeTime:     zapcore.ISO8601TimeEncoder,       // 时间编码器
		EncodeDuration: zapcore.SecondsDurationEncoder,   //
		EncodeCaller:   zapcore.ShortCallerEncoder,       //
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 这里可以指定颜色
	}

	if err := os.MkdirAll(param.LogDir, os.ModePerm); err != nil {
		panic("无法创建日志文件夹: " + err.Error())
	}

	logPath := param.LogDir + "/app.log"

	var cores []zapcore.Core

	// 控制台输出
	if param.ConsoleOutPut {
		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		consoleCore := zapcore.NewCore(
			consoleEncoder,
			zapcore.Lock(os.Stdout),
			logLevel,
		)
		cores = append(cores, consoleCore)
	}

	// 文件输出
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    100,  // 单个日志文件的最大大小/MB
		MaxBackups: 3,    // 指定要保留的旧日志文件的最大数量
		MaxAge:     1,    // 按天切割
		LocalTime:  true, // 使用本地时间
		Compress:   true, // 压缩历史日志
	}

	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	fileCore := zapcore.NewCore(
		fileEncoder,
		zapcore.AddSync(lumberjackLogger),
		logLevel,
	)
	cores = append(cores, fileCore)

	core := zapcore.NewTee(cores...)

	// 创建日志记录器
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.WarnLevel))
}

// WithContext 添加 requestid 到 logger
func WithContext(ctx context.Context) zap.Option {
	if Logger == nil {
		panic("Logger not init")
	}
	requestID, ok := ctx.Value(RequestIDKey).(string)
	if !ok {
		requestID = "unknown-request-id"
	}

	return zap.Fields(zap.String(RequestIDKey, requestID))
}
