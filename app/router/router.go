package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yourusername/fiber-template/app/controller"
	"github.com/yourusername/fiber-template/app/middleware"
)

// SetupRoutes 设置应用程序的所有路由
func SetupRoutes(app *fiber.App) {
	// 配置静态资源
	app.Static("/public", "./app/public")

	// API 分组
	api := app.Group("/api")

	// 健康检查
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// v1 版本分组
	v1 := api.Group("/v1")

	// 认证路由
	auth := v1.Group("/auth")
	auth.Post("/login", controller.Login)
	auth.Post("/register", controller.Register)

	// 用户路由 - 需要认证
	userRoutes := v1.Group("/users")
	userRoutes.Get("/", controller.GetUsers)
	userRoutes.Get("/:id", controller.GetUser)
	userRoutes.Post("/", controller.CreateUser, middleware.AuthRequired())
	userRoutes.Put("/:id", controller.UpdateUser, middleware.AuthRequired())
	userRoutes.Delete("/:id", controller.DeleteUser, middleware.AuthRequired())
}
