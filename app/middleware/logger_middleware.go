package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupLogger 设置Fiber的日志中间件
func SetupLogger(app *fiber.App) {
	// 创建日志文件
	file, err := os.OpenFile("./logs/fiber.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		// 如果创建文件失败，将日志输出到控制台
		app.Use(logger.New(logger.Config{
			Format:     "${time} | ${status} | ${latency} | ${method} | ${path}\n",
			TimeFormat: "2006-01-02 15:04:05",
		}))
		return
	}

	// 配置Fiber的日志中间件
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${status} | ${latency} | ${method} | ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		Output:     file,
	}))
}
