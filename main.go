package main

import (
	"github.com/gofiber/fiber/v3"

	. "autopub-server/router"
	. "autopub-server/util"
)

func main() {
	// 创建一个Fiber 应用实例
	app := App()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendFile("index.html")
	})

	// 自动注册路由
	Router()

	// 如果监听失败，则输出错误信息并终止程序
	if err := app.Listen("0.0.0.0:8083"); err != nil {
		panic(err)
	}
}
