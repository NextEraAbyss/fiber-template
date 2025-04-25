package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/yourusername/fiber-template/app/model"
	"github.com/yourusername/fiber-template/app/router"
	"github.com/yourusername/fiber-template/app/schedule"
	"github.com/yourusername/fiber-template/config"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 初始化数据库连接
	config.InitDB(cfg)

	// 自动迁移数据库模型
	db := config.GetDB()
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("无法迁移数据库模型: %v", err)
	}

	// 初始化并启动定时任务
	schedule.InitTasks()
	schedule.BeginTasks()

	// 创建新的 Fiber 实例
	app := fiber.New(fiber.Config{
		AppName:      cfg.AppName,
		ErrorHandler: config.ErrorHandler,
	})

	// 中间件
	app.Use(recover.New())
	app.Use(logger.New())

	// 设置路由
	router.SetupRoutes(app)

	// 启动服务器
	log.Fatal(app.Listen(":" + cfg.Port))
}
