package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/yourusername/fiber-template/config"
)

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

		// 获取配置
		cfg := config.Load()

		// 验证令牌
		claims, err := config.ValidateToken(tokenParts[1], cfg)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "无效的令牌: " + err.Error(),
			})
		}

		// 将用户信息保存到上下文
		c.Locals("user_id", claims.UserID)
		c.Locals("email", claims.Email)

		// 继续下一个中间件/处理程序
		return c.Next()
	}
}
