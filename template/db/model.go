package db

import (
	"context"

	"github.com/bigbigliu/go_core/database/mysql"
	"github.com/bigbigliu/go_core/pkgs"
	"gorm.io/gorm"
)

type IAppNew interface {
	// Create 创建app
	Create(ctx context.Context, param *App) error
	// GetList 查询列表
	GetList(ctx context.Context, param *GetListQueryParam) (count int64, rows []*App, err error)
	// Detail 查询详情
	Detail(ctx context.Context, param *App) (row *App, err error)
	// Edit 编辑
	Edit(ctx context.Context, param *App) error
	// Delete 编辑
	Delete(ctx context.Context, param *App) error
}

// App 第三方app
type App struct {
	mysql.BasicModel
	AppID       string `gorm:"column:app_id" json:"app_id"`               // AppID app_id
	AppSecret   string `gorm:"column:app_secret" json:"app_secret"`       // AppSecret app密码
	AppName     string `gorm:"column:app_name" json:"app_name"`           // AppName app名称
	AppHomePage string `gorm:"column:app_home_page" json:"app_home_page"` // AppHomePage app主页url
}

// GetListQueryParam 列表查询参数
type GetListQueryParam struct {
	QueryParam  pkgs.ReqQuery
	AppName     string `gorm:"column:app_name" json:"app_name"`           // app名称
	AppHomePage string `gorm:"column:app_home_page" json:"app_home_page"` // app主页url
}

// GetTable 获取表名
func (c *App) GetTable() *gorm.DB {
	return mysql.DBClient.Table("app_info")
}

// Create 创建app
func (c *App) Create(ctx context.Context, param *App) error {
	err := c.GetTable().WithContext(ctx).Create(&param).Error
	return err
}

// GetList 查询列表
func (c *App) GetList(ctx context.Context, param *GetListQueryParam) (count int64, rows []*App, err error) {
	tx := c.GetTable().WithContext(ctx)

	if param.AppName != "" {
		tx.Where("app_name like ?", param.AppName+"%")
	}

	if param.AppHomePage != "" {
		tx.Where("app_home_page = ?", param.AppHomePage)
	}

	tx.Count(&count)

	err = tx.Order("id desc").Limit(param.QueryParam.Limit).Offset(param.QueryParam.Offset).Find(&rows).Error
	return count, rows, err
}

// Detail 查询详情
func (c *App) Detail(ctx context.Context, param *App) (row *App, err error) {
	err = c.GetTable().WithContext(ctx).Where(param).First(&row).Error
	return row, err
}

// Edit 编辑
func (c *App) Edit(ctx context.Context, param *App) error {
	err := c.GetTable().WithContext(ctx).Updates(param).Error
	return err
}

// Delete 编辑
func (c *App) Delete(ctx context.Context, param *App) error {
	err := c.GetTable().WithContext(ctx).Delete(param).Error
	return err
}
