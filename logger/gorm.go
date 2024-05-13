package logger

import (
	"context"
	"runtime"
	"time"

	"go.uber.org/zap"
	gormLog "gorm.io/gorm/logger"
)

// CustomLogger gorm日志记录器
type CustomLogger struct {
	logger  *zap.Logger
	context context.Context
	level   gormLog.LogLevel
}

// NewCustomLogger gorm日志记录器
func NewCustomLogger(zapLogger *zap.Logger, ctx context.Context, level gormLog.LogLevel) *CustomLogger {
	return &CustomLogger{
		logger:  zapLogger,
		context: ctx,
		level:   level, // 设置日志级别
	}
}

// LogMode 设置日志记录模式
func (l *CustomLogger) LogMode(level gormLog.LogLevel) gormLog.Interface {
	l.level = level
	return l
}

// Info 记录信息日志
func (l *CustomLogger) Info(ctx context.Context, s string, args ...interface{}) {
	if l.level >= gormLog.Info {
		l.logger.Info(s, append([]zap.Field{zap.Any(RequestIDKey, ctx.Value(RequestIDKey))}, zap.Any("arguments", args))...)
	}
}

// Warn 记录警告日志
func (l *CustomLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	if l.level >= gormLog.Warn {
		l.logger.Warn(s, append([]zap.Field{zap.Any(RequestIDKey, ctx.Value(RequestIDKey))}, zap.Any("arguments", args))...)
	}
}

// Error 记录错误日志
func (l *CustomLogger) Error(ctx context.Context, s string, args ...interface{}) {
	if l.level >= gormLog.Error {
		l.logger.Error(s, append([]zap.Field{zap.Any(RequestIDKey, ctx.Value(RequestIDKey))}, zap.Any("arguments", args))...)
	}
}

// Trace 记录追踪日志
func (l *CustomLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.level == gormLog.Silent {
		return
	}

	funcName, funcline := getCallingFunction()

	s, ts := fc()
	fields := []zap.Field{
		zap.Any(RequestIDKey, ctx.Value(RequestIDKey)), // requestID
		zap.String("sql", s),                           // sql语句
		zap.Int64("rows", ts),                          // 受影响行数
		zap.String("function", funcName),               // 记录执行的 SQL 函数
		zap.Int("function_line", funcline),             // 记录执行的 SQL 函数行号
		zap.Duration("elapsed", time.Since(begin)),     // sql耗时 / 纳秒
	}
	if err != nil {
		if err.Error() != "record not found" {
			l.logger.Error("SQL Error", append(fields, zap.Error(err))...)
		}
	} else {
		l.logger.Info("SQL Query", fields...)
	}
}

// getCallingFunction 获取当前调用的函数名
func getCallingFunction() (string, int) {
	pc := make([]uintptr, 10)
	runtime.Callers(5, pc)
	frames := runtime.CallersFrames(pc)
	frame, _ := frames.Next()
	return frame.Function, frame.Line
}
