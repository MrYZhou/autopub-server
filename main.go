package main

import (
	"github.com/gofiber/fiber/v2"

	. "autopub-server/api"
)





func main() {
	// 创建一个新的 Fiber 应用实例
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("autopub server")
	})
	
	app.Get("pubweb",func(c *fiber.Ctx) error {
		Pubweb("D:/Users/JNPF/Desktop/jnpf-crm2-web","/root/testweb")
		return c.SendString("ok")
	})

	app.Post("pubjava",func(c *fiber.Ctx) error {
		var model JarUpload
		// 从请求体中读取JSON内容并反序列化
		if err := c.BodyParser(&model); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON body")
		}
		err:=Pubjava(model)
		model.Msg = "success"
		if err!=nil{
			model.Msg = err.Error()
		}
		return c.JSON(model)
	})

	// 设置服务器监听地址和端口
	if err := app.Listen(":8083"); err != nil {
		// 如果监听失败，则输出错误信息并终止程序
		panic(err)
	}
}
