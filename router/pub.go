package router

import (
	"github.com/gofiber/fiber/v2"

	. "autopub-server/server"
	. "autopub-server/util"
)

func init(){
	app := App()

	// 创建子路由
	api := app.Group("/pub")
	api.Get("/a",func (c *fiber.Ctx) error {
		return AppResult(c).Success()
	} )

	app.Post("pubweb", Handlepubweb)

	app.Post("pubjava", Handlepubjava)
}


func Handlepubweb(c *fiber.Ctx) error {
	var model WebrUpload
	// 从请求体中读取JSON内容并反序列化
	if err := c.BodyParser(&model); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON body")
	}

	err := Pubweb(model)
	if err != nil {
		return AppResult(c).Fail(err.Error())
	}
	return AppResult(c).Success(model,"部署web完成")
}

func Handlepubjava(c *fiber.Ctx) error {
	var model JarUpload
	// 从请求体中读取JSON内容并反序列化
	if err := c.BodyParser(&model); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON body")
	}
	err := Pubjava(model)
	if err != nil {
		return AppResult(c).Fail(err.Error())
	}
	return AppResult(c).Success(model,"部署java完成")

}
