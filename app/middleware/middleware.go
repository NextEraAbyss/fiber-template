package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Setup 配置所有中间件
func Setup(app *fiber.App) {
	// 错误恢复中间件
	app.Use(recover.New())

	// 安全相关中间件
	SetupSecurity(app)

	// 日志中间件
	SetupLogger(app)
}
