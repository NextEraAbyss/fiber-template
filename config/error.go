package config

import (
	"github.com/gofiber/fiber/v2"
)

// ErrorHandler 是应用程序的全局错误处理器
func ErrorHandler(c *fiber.Ctx, err error) error {
	// 默认 500 状态码
	code := fiber.StatusInternalServerError

	// 检查是否是 Fiber 错误
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// 返回 JSON 响应
	return c.Status(code).JSON(fiber.Map{
		"success": false,
		"error":   err.Error(),
	})
}
