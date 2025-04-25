package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/NextEraAbyss/fiber-template/app/model"
	"github.com/NextEraAbyss/fiber-template/config"
)

// GetAllUsers 返回所有用户
func GetAllUsers() ([]*model.User, error) {
	var users []*model.User
	result := config.GetDB().Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// GetUserByID 通过ID返回用户
func GetUserByID(id string) (*model.User, error) {
	var user model.User

	// 将字符串ID转换为uint
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, errors.New("无效的用户ID")
	}

	result := config.GetDB().First(&user, uint(userID))
	if result.Error != nil {
		return nil, errors.New("用户未找到")
	}
	return &user, nil
}

// GetUserByEmail 通过邮箱查找用户
func GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	result := config.GetDB().Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, errors.New("用户未找到")
	}

	return &user, nil
}

// CreateUser 创建新用户
func CreateUser(user *model.User) (*model.User, error) {
	// 设置创建和更新时间戳
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	// 保存用户到数据库
	result := config.GetDB().Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// UpdateUser 通过ID更新用户
func UpdateUser(id string, userData *model.User) (*model.User, error) {
	var user model.User

	// 将字符串ID转换为uint
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, errors.New("无效的用户ID")
	}

	// 查询用户是否存在
	result := config.GetDB().First(&user, uint(userID))
	if result.Error != nil {
		return nil, errors.New("用户未找到")
	}

	// 更新用户字段
	user.Username = userData.Username
	user.Email = userData.Email
	if userData.Password != "" {
		user.Password = userData.Password
	}
	user.UpdatedAt = time.Now()

	// 保存更新
	config.GetDB().Save(&user)

	return &user, nil
}

// DeleteUser 通过ID删除用户
func DeleteUser(id string) error {
	// 将字符串ID转换为uint
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return errors.New("无效的用户ID")
	}

	// 软删除用户
	result := config.GetDB().Delete(&model.User{}, uint(userID))
	if result.Error != nil {
		return result.Error
	}

	// 检查是否找到并删除了用户
	if result.RowsAffected == 0 {
		return errors.New("用户未找到")
	}

	return nil
}
