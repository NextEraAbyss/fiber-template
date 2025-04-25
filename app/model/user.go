package model

import (
	"time"

	"github.com/NextEraAbyss/fiber-template/config"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 表示系统中的用户
// @Description 用户信息模型
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey" example:"1"`
	Username  string         `json:"username" gorm:"size:100;not null;unique" example:"john_doe" validate:"required,min=3,max=32,safename"`
	Email     string         `json:"email" gorm:"size:100;not null;unique" example:"john@example.com" validate:"required,email"`
	Password  string         `json:"-" gorm:"size:100;not null" example:"password123" validate:"required,min=8,max=100"`
	Avatar    string         `json:"avatar" gorm:"size:255" example:"https://example.com/avatar.jpg"`
	Role      string         `json:"role" gorm:"size:20;default:user" example:"user" validate:"required,oneof=user admin"`
	IsActive  bool           `json:"is_active" gorm:"default:true" example:"true"`
	LastLogin *time.Time     `json:"last_login,omitempty" example:"2024-01-01T00:00:00Z"`
	CreatedAt time.Time      `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt time.Time      `json:"updated_at" example:"2024-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

// BeforeCreate 创建前钩子
func (u *User) BeforeCreate(tx *gorm.DB) error {
	// 如果密码未被哈希处理，则进行哈希处理
	if len(u.Password) < 60 { // bcrypt哈希通常长度为60
		// 对密码进行安全转义处理
		password := config.SanitizeString(u.Password)

		// 使用bcrypt进行哈希
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}

	// 清理和验证用户名
	u.Username = config.SanitizeString(u.Username)

	// 设置默认角色
	if u.Role == "" {
		u.Role = "user"
	}

	return nil
}

// Validate 验证用户数据
func (u *User) Validate() []*config.ValidationError {
	return config.ValidateStruct(u)
}
