package main

import (
	. "autopub-server/router"
	. "autopub-server/util"
	"embed"
	"os"

	"github.com/gofiber/fiber/v3"
)

//go:embed .env
var embedDir embed.FS

//go:embed resources/**
var embedResources embed.FS

func main() {
	// 创建一个Fiber 应用实例
	app := App()

	app.Get("/", func(c fiber.Ctx) error {
		return c.Redirect().To("https://larryer.xyz/")
	})

	// 自动注册路由
	Router()

	// 如果监听失败，则输出错误信息并终止程序
	if err := app.Listen(os.Getenv("serverHost")); err != nil {
		panic(err)
	}
}
