package main

import (
	"os"

	"github.com/gofiber/fiber/v2"

	. "autopub-server/router"
	. "autopub-server/server"
)

func main() {
	// 创建一个Fiber 应用实例
	app := fiber.New()
	// 注册自定义中间件以转换上下文
	app.Use(CtxMiddleware)

	app.Get("/", func(c *fiber.Ctx) error {
		return AppResult(c).Success("autopub server")
	})

	app.Post("pubweb", Handlepubweb)

	app.Post("pubjava", Handlepubjava)

	// 设置服务器地址
	mode := os.Getenv("MODE")
	url := "127.0.0.1:8083"
	if mode == "production" {
		url = "0.0.0.0:8083"
	}
	if err := app.Listen(url); err != nil {
		// 如果监听失败，则输出错误信息并终止程序
		panic(err)
	}
}
