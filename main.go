package main

import (
	"github.com/gofiber/fiber/v2"

	. "autopub-server/router"
	. "autopub-server/server"
)


func main() {
	// 创建一个新的 Fiber 应用实例
	app := fiber.New()
	// 注册自定义中间件以转换上下文
	app.Use(CtxMiddleware)

	app.Get("/", func(c *fiber.Ctx) error {
		return Success(c,"autopub server")
	})
	
	app.Post("pubweb",Handlepubweb)

	app.Post("pubjava",Handlepubjava)

	// 设置服务器监听地址和端口
	if err := app.Listen(":8083"); err != nil {
		// 如果监听失败，则输出错误信息并终止程序
		panic(err)
	}
}
