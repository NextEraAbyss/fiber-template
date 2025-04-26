package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// 安全相关配置
var (
	// CORS配置
	corsConfig = cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: false,
		MaxAge:           300,
	}

	// 速率限制配置
	limiterConfig = limiter.Config{
		Max:        100,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"status":  "error",
				"message": "Too many requests",
			})
		},
	}
)

// SetupSecurity 设置安全相关的中间件
func SetupSecurity(app *fiber.App) {
	// 设置Helmet中间件，用于保护常见Web漏洞
	app.Use(helmet.New())

	// 设置CORS中间件
	app.Use(cors.New(corsConfig))

	// 设置速率限制中间件，防止暴力攻击
	app.Use(limiter.New(limiterConfig))
}
