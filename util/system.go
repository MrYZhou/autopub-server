package util

import (
	"os"

	"github.com/bytedance/sonic"
	"github.com/gofiber/contrib/monitor"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

var app *fiber.App

func App() *fiber.App {
	if app == nil {
		app = fiber.New(fiber.Config{
			JSONEncoder: sonic.Marshal,
			JSONDecoder: sonic.Unmarshal,
		})
		// 注册自定义中间件以转换上下文
		app.Use(CtxMiddleware)
		// 静态文件服务
		app.Get("/file*", static.New(os.Getenv("resources")))
		// 监控
		app.Use("/metrics", monitor.New())
	}
	return app
}
