package middleware

import (
	"errors"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// TokenMetadata 包含JWT令牌的元数据
type TokenMetadata struct {
	UserID  uint
	Expires time.Time
}

// TokenClaims 自定义JWT声明结构
type TokenClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

// SECRET_KEY JWT密钥(应从环境变量/配置中获取)
const SECRET_KEY = "your_secret_key"

// AuthRequired 是一个检查用户是否已认证的中间件
func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 从header获取令牌
		authHeader := c.Get("Authorization")

		// 检查是否存在Authorization头
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "需要认证",
			})
		}

		// 提取令牌（通常格式为"Bearer token"）
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "无效的认证格式",
			})
		}

		// 验证令牌
		token, err := verifyToken(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "无效的令牌: " + err.Error(),
			})
		}

		// 获取Claims
		claims, ok := token.Claims.(*TokenClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "无法解析token claims",
			})
		}

		// 将用户信息保存到上下文
		c.Locals("user_id", claims.UserID)

		// 继续下一个中间件/处理程序
		return c.Next()
	}
}

// ExtractTokenMetadata 从请求中提取TOKEN元数据
func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	// 设置默认值
	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return nil, errors.New("无法解析token claims")
	}

	// 检查token是否过期
	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token已过期")
	}

	return &TokenMetadata{
		UserID:  claims.UserID,
		Expires: time.Unix(claims.ExpiresAt, 0),
	}, nil
}

// extractToken 从请求中提取token
func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	// 通常，授权头的格式为"Bearer token"
	// 检查是否提供了Bearer
	parts := strings.Split(bearToken, " ")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}

// verifyToken 验证token
func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 确保使用期望的签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// CreateToken 创建一个新token
func CreateToken(userID uint) (string, error) {
	// 设置token有效期为24小时
	expirationTime := time.Now().Add(24 * time.Hour)

	// 创建JWT声明，包括用户ID和过期时间
	claims := &TokenClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	// 使用claims创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名token并获取完整的编码token作为字符串
	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Protected 保护路由，只允许已认证的用户访问
func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		metadata, err := ExtractTokenMetadata(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "未授权访问",
				"error":   err.Error(),
			})
		}

		// 将用户ID存储在上下文中，以便后续处理函数使用
		c.Locals("user_id", metadata.UserID)

		return c.Next()
	}
}
