package router

import (
	"github.com/NextEraAbyss/fiber-template/app/controller"
	"github.com/NextEraAbyss/fiber-template/app/middleware"
	"github.com/gofiber/fiber/v2"
)

// SetupAuthRoutes 设置认证相关路由
func SetupAuthRoutes(router fiber.Router) {
	// 创建认证控制器
	authController := controller.NewAuthController()

	// 认证路由组
	auth := router.Group("/auth")

	// 公开路由
	auth.Post("/register", authController.Register)
	auth.Post("/login", authController.Login)

	// 需要认证的路由
	auth.Get("/me", middleware.Protected(), authController.GetMe)
}
