package router

import (
	"github.com/NextEraAbyss/fiber-template/app/controller"
	"github.com/NextEraAbyss/fiber-template/app/middleware"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes 设置应用程序的所有路由
func SetupRoutes(app *fiber.App) {
	// 配置静态资源
	app.Static("/public", "./app/public")

	// 设置全局中间件
	middleware.Setup(app)

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

	// 初始化控制器
	userController := controller.NewUserController()

	// 用户路由
	v1.Get("/users", userController.GetUsers)    // 获取用户列表
	v1.Get("/users/:id", userController.GetUser) // 获取单个用户
}
