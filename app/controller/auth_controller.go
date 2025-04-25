package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yourusername/fiber-template/app/model"
	"github.com/yourusername/fiber-template/app/service"
	"github.com/yourusername/fiber-template/config"
	"golang.org/x/crypto/bcrypt"
)

// LoginRequest 登录请求体
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterRequest 注册请求体
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login 处理用户登录
func Login(c *fiber.Ctx) error {
	// 解析请求体
	var loginReq LoginRequest
	if err := c.BodyParser(&loginReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "无效的请求体",
		})
	}

	// 通过邮箱查找用户
	user, err := service.GetUserByEmail(loginReq.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "用户名或密码错误",
		})
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "用户名或密码错误",
		})
	}

	// 生成JWT令牌
	cfg := config.Load()
	token, err := config.GenerateToken(user.ID, user.Email, cfg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "生成令牌失败",
		})
	}

	// 返回认证成功响应
	return c.JSON(fiber.Map{
		"success": true,
		"message": "登录成功",
		"token":   token,
		"user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

// Register 处理用户注册
func Register(c *fiber.Ctx) error {
	// 解析请求体
	var registerReq RegisterRequest
	if err := c.BodyParser(&registerReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "无效的请求体",
		})
	}

	// 检查邮箱是否已被使用
	_, err := service.GetUserByEmail(registerReq.Email)
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "邮箱已被注册",
		})
	}

	// 哈希密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "密码加密失败",
		})
	}

	// 创建新用户
	user := &model.User{
		Username: registerReq.Username,
		Email:    registerReq.Email,
		Password: string(hashedPassword),
	}

	// 保存用户
	newUser, err := service.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	// 生成JWT令牌
	cfg := config.Load()
	token, err := config.GenerateToken(newUser.ID, newUser.Email, cfg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "生成令牌失败",
		})
	}

	// 返回注册成功响应
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "注册成功",
		"token":   token,
		"user": fiber.Map{
			"id":       newUser.ID,
			"username": newUser.Username,
			"email":    newUser.Email,
		},
	})
}
