package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 用户角色常量
const (
	RoleAdmin = "admin"
	RoleUser  = "user"
	RoleGuest = "guest"
)

// 用户状态常量
const (
	UserActive   = 1
	UserInactive = 0
)

// User 表示系统中的用户
// @Description 用户信息模型
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"size:100;not null;unique" validate:"required,min=3,max=100"`
	Email    string `json:"email" gorm:"size:100;not null;unique" validate:"required,email"`
	Password string `json:"-" gorm:"size:100;not null" validate:"required,min=6"`
	Avatar   string `json:"avatar" gorm:"size:255"`
	Role     string `json:"role" gorm:"size:20;default:user" validate:"oneof=admin user guest"`
	IsActive int    `json:"is_active" gorm:"default:1" validate:"oneof=0 1"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

// BeforeCreate 创建前钩子
func (u *User) BeforeCreate(tx *gorm.DB) error {
	return u.hashPasswordIfNeeded()
}

// BeforeUpdate 更新前钩子
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	return u.hashPasswordIfNeeded()
}

// hashPasswordIfNeeded 如果密码未哈希则进行哈希处理
func (u *User) hashPasswordIfNeeded() error {
	// 如果密码未被哈希处理，则进行哈希处理
	if len(u.Password) < 60 { // bcrypt哈希通常长度为60
		// 使用bcrypt进行哈希
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}

	// 设置默认角色
	if u.Role == "" {
		u.Role = RoleUser
	}

	// 设置默认激活状态
	if u.IsActive != UserActive && u.IsActive != UserInactive {
		u.IsActive = UserActive
	}

	return nil
}

// CheckPassword 验证密码是否正确
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// ChangePassword 更改用户密码
func (u *User) ChangePassword(newPassword string) error {
	u.Password = newPassword
	return u.hashPasswordIfNeeded()
}

// SanitizeOutput 清理敏感数据用于输出
func (u *User) SanitizeOutput() map[string]interface{} {
	return map[string]interface{}{
		"id":         u.ID,
		"username":   u.Username,
		"email":      u.Email,
		"avatar":     u.Avatar,
		"role":       u.Role,
		"is_active":  u.IsActive,
		"created_at": u.CreatedAt,
		"updated_at": u.UpdatedAt,
	}
}
