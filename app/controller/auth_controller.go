package controller

import (
	"time"

	"github.com/NextEraAbyss/fiber-template/app/middleware"
	"github.com/NextEraAbyss/fiber-template/app/model"
	"github.com/NextEraAbyss/fiber-template/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// AuthController 处理用户认证相关请求
type AuthController struct {
	DB *gorm.DB
}

// NewAuthController 创建新的AuthController实例
func NewAuthController() *AuthController {
	return &AuthController{
		DB: config.GetDB(),
	}
}

// RegisterPayload 用户注册请求有效载荷
type RegisterPayload struct {
	Username string `json:"username" validate:"required,min=3,max=32,safename"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

// LoginPayload 用户登录请求有效载荷
type LoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Register 处理用户注册请求
func (c *AuthController) Register(ctx *fiber.Ctx) error {
	var user model.User
	if err := ctx.BodyParser(&user); err != nil {
		return config.BadRequestError(ctx, "无效的请求数据")
	}

	// 验证用户数据
	if errors := config.ValidateStruct(user); errors != nil {
		return config.ValidationErrors(ctx, "验证失败")
	}

	// 检查用户名是否已存在
	var existingUser model.User
	if err := c.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return config.ConflictError(ctx, "用户名已存在")
	}

	// 检查邮箱是否已存在
	if err := c.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return config.ConflictError(ctx, "邮箱已被注册")
	}

	// 创建用户
	if err := c.DB.Create(&user).Error; err != nil {
		return config.ServerError(ctx)
	}

	// 生成 token
	token, err := middleware.CreateToken(user.ID)
	if err != nil {
		return config.ServerError(ctx)
	}

	return config.Success(ctx, fiber.Map{
		"token": token,
		"user": fiber.Map{
			"id":        user.ID,
			"username":  user.Username,
			"email":     user.Email,
			"avatar":    user.Avatar,
			"role":      user.Role,
			"isActive":  user.IsActive,
			"lastLogin": user.LastLogin,
		},
	})
}

// Login 处理用户登录请求
func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var loginData struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	if err := ctx.BodyParser(&loginData); err != nil {
		return config.BadRequestError(ctx, "无效的请求数据")
	}

	// 验证登录数据
	if errors := config.ValidateStruct(loginData); errors != nil {
		return config.ValidationErrors(ctx, "验证失败")
	}

	var user model.User
	if err := c.DB.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
		return config.UnauthorizedError(ctx)
	}

	// 验证密码
	if !config.CheckPasswordHash(loginData.Password, user.Password) {
		return config.UnauthorizedError(ctx)
	}

	// 更新最后登录时间
	now := time.Now()
	user.LastLogin = &now
	c.DB.Save(&user)

	// 生成 token
	token, err := middleware.CreateToken(user.ID)
	if err != nil {
		return config.ServerError(ctx)
	}

	return config.Success(ctx, fiber.Map{
		"token": token,
		"user": fiber.Map{
			"id":        user.ID,
			"username":  user.Username,
			"email":     user.Email,
			"avatar":    user.Avatar,
			"role":      user.Role,
			"isActive":  user.IsActive,
			"lastLogin": user.LastLogin,
		},
	})
}

// GetMe 获取当前已认证用户的信息
func (c *AuthController) GetMe(ctx *fiber.Ctx) error {
	// 从上下文中获取用户ID
	userID := ctx.Locals("user_id").(uint)

	// 查找用户
	var user model.User
	if result := c.DB.First(&user, userID); result.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "未找到用户",
		})
	}

	// 返回用户信息
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "获取用户信息成功",
		"data": fiber.Map{
			"user": fiber.Map{
				"id":         user.ID,
				"username":   user.Username,
				"email":      user.Email,
				"avatar":     user.Avatar,
				"role":       user.Role,
				"is_active":  user.IsActive,
				"last_login": user.LastLogin,
				"created_at": user.CreatedAt,
				"updated_at": user.UpdatedAt,
			},
		},
	})
}
