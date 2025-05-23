package main

import (
	"log"

	"github.com/NextEraAbyss/fiber-template/app/model"
	"github.com/NextEraAbyss/fiber-template/app/router"
	"github.com/NextEraAbyss/fiber-template/config"
	_ "github.com/NextEraAbyss/fiber-template/docs" // swagger文档
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

//	@title			Fiber Template API
//	@version		1.0
//	@description	Fiber框架模板API文档
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.example.com/support
//	@contact.email	support@example.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:3000
// @BasePath	/api/v1
// @schemes	http
func main() {
	// 加载配置
	cfg := config.Load()

	// 初始化应用
	initApp(cfg)

	// 创建Fiber实例
	app := fiber.New(fiber.Config{
		AppName:      cfg.App.Name,
		ErrorHandler: config.ErrorHandler,
	})

	// Swagger路由
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// 设置路由
	router.SetupRoutes(app)

	// 启动服务器
	log.Fatal(app.Listen(":" + cfg.App.Port))
}

// initApp 初始化应用程序
func initApp(cfg *config.Config) {
	// 初始化数据库连接
	config.InitDB(cfg)

	// 自动迁移数据库模型
	db := config.GetDB()
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("无法迁移数据库模型: %v", err)
	}

	// 初始化并启动定时任务
	config.InitTasks()
	config.BeginTasks()
}
