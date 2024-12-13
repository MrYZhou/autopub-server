package util

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func App() *fiber.App {
	if app == nil {
		app = fiber.New()
		// 注册自定义中间件以转换上下文
		app.Use(CtxMiddleware)
		// 静态文件服务
		app.Static("/file", os.Getenv("resources"))
	}
	return app
}
