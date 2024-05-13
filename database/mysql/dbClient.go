package mysql

import (
	"context"
	"fmt"
	"os"

	"github.com/bigbigliu/go_core/logger"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
)

var (
	// DBClient 全局DB客户端
	DBClient *gorm.DB
	err      error
)

// GenerateDSNParam 生成dsn参数
type GenerateDSNParam struct {
	DbHost string `json:"db_host"` // DbHost 数据库服务host
	DbPort int    `json:"db_port"` // DbPort 数据库服务port
	DbUser string `json:"db_user"` // DbUser 数据库服务用户名
	DbPwd  string `json:"db_pwd"`  // DbPwd 数据库服务密码
	DbName string `json:"db_name"` // DbName 数据库名
}

// InitDB 初始化DB连接
func (h *GenerateDSNParam) InitDB() {
	logger.Logger.Info("DB", zap.String("conn", "connecting..."))
	DBClient, err = gorm.Open(mysql.Open(
		h.GenerateDSN()),
		&gorm.Config{
			Logger: logger.NewCustomLogger(logger.Logger, context.Background(), gormLog.Info),
		})
	if err != nil {
		logger.Logger.Info("InitDB Error: ", zap.Error(err))
		os.Exit(-1)
	}
	logger.Logger.Info("DB", zap.String("conn", "数据库连接成功"))
	sqlDB, _ := DBClient.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}

// GenerateDSN dsn
func (h *GenerateDSNParam) GenerateDSN() string {
	user := h.DbUser
	pwd := h.DbPwd
	dbmame := h.DbName
	host := h.DbHost
	port := h.DbPort

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pwd, host, port, dbmame)
	logger.Logger.Info("DB", zap.String("dsn", dsn))
	return dsn
}
