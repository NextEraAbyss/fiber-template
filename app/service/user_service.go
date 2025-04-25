package service

import (
	"errors"
	"strconv"

	"github.com/NextEraAbyss/fiber-template/app/model"
	"github.com/NextEraAbyss/fiber-template/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// 定义错误
var (
	ErrUserNotFound = errors.New("用户不存在")
)

// UserQueryParams 用户查询参数
type UserQueryParams struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	IsActive string `json:"is_active"`
}

// UserService 用户服务
type UserService struct {
	db *gorm.DB
}

// NewUserService 创建新的用户服务实例
func NewUserService() *UserService {
	return &UserService{
		db: config.GetDB(),
	}
}

// GetUsers 获取用户列表
func (s *UserService) GetUsers(params UserQueryParams) ([]model.User, fiber.Map, error) {
	// 构建查询
	query := s.db.Model(&model.User{})

	// 添加查询条件
	if params.Username != "" {
		query = query.Where("username LIKE ?", "%"+params.Username+"%")
	}
	if params.Email != "" {
		query = query.Where("email LIKE ?", "%"+params.Email+"%")
	}
	if params.Role != "" {
		query = query.Where("role = ?", params.Role)
	}
	if params.IsActive != "" {
		// 将字符串转换为整数
		active, err := strconv.Atoi(params.IsActive)
		if err == nil {
			query = query.Where("is_active = ?", active)
		}
	}

	// 计算总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, nil, err
	}

	// 计算总页数
	totalPages := (total + int64(params.PageSize) - 1) / int64(params.PageSize)

	// 获取分页数据
	var users []model.User
	if err := query.Offset((params.Page - 1) * params.PageSize).Limit(params.PageSize).Find(&users).Error; err != nil {
		return nil, nil, err
	}

	// 构建分页信息
	pagination := fiber.Map{
		"total":       total,
		"page":        params.Page,
		"page_size":   params.PageSize,
		"total_pages": totalPages,
	}

	return users, pagination, nil
}

// GetUserByID 通过ID获取用户
func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := s.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
