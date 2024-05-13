package mysql

import (
	"gorm.io/gorm"
)

// BasicModel 基础模型
type BasicModel struct {
	gorm.Model
	CreatedUser string `gorm:"column:created_user" json:"created_user"`
	UpdatedUser string `gorm:"column:updated_user" json:"updated_user"`
	DeletedUser string `gorm:"column:deleted_user" json:"deleted_user"`
}
