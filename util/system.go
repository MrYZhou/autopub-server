package util

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func init() {
	dbUrl:= os.Getenv("dbUrl")
	if dbUrl == "" {
		dbUrl = "root:root@tcp(127.0.0.1:3306)/study" 
	}
	DbInit("default",dbUrl)
	DbChange("default")
	
}

var app  *fiber.App

func App() *fiber.App{
	if app == nil{
		app = fiber.New()
		// 注册自定义中间件以转换上下文
		app.Use(CtxMiddleware)
		// 静态文件服务
		app.Static("/file", "./static")		
	}
	return app
}