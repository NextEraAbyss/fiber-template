package router

import (
	"github.com/gofiber/fiber/v2"
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

	// 设置各个模块的路由
	SetupAuthRoutes(v1)
	SetupUserRoutes(v1)

	// 设置中间件
	middleware.SetupSecurity(app)
	middleware.SetupLogger(app)
}

// SetupUserRoutes 设置用户相关路由
func SetupUserRoutes(router fiber.Router) {
	// 用户路由组 - 需要认证
	// 当添加用户控制器后取消注释
	/*
		userRoutes := router.Group("/users", middleware.Protected())

		// 这里添加用户相关路由
		// userRoutes.Get("/", userController.GetUsers)
		// userRoutes.Get("/:id", userController.GetUser)
		// userRoutes.Post("/", userController.CreateUser)
		// userRoutes.Put("/:id", userController.UpdateUser)
		// userRoutes.Delete("/:id", userController.DeleteUser)
	*/
}
