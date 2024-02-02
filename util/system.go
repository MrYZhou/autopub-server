package util

import (

	"github.com/gofiber/fiber/v2"
)

func Init() {
	DbInit()
}

var app  *fiber.App

func App() *fiber.App{
	if app == nil{
		app = fiber.New()
		// 注册自定义中间件以转换上下文
		app.Use(CtxMiddleware)
	}
	return app
}