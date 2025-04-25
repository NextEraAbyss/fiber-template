package model

import (
	"time"

	"gorm.io/gorm"
)

// User 表示系统中的用户
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"size:100;not null;unique"`
	Email     string         `json:"email" gorm:"size:100;not null;unique"`
	Password  string         `json:"-" gorm:"size:100;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

// BeforeCreate 创建前钩子
func (u *User) BeforeCreate(tx *gorm.DB) error {
	// 这里可以添加密码哈希等逻辑
	return nil
}
