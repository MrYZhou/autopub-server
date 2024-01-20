package router

import (
	"github.com/gofiber/fiber/v2"

	. "autopub-server/server"
)

func Handlepubweb(c *fiber.Ctx) error {
	var model WebrUpload
	// 从请求体中读取JSON内容并反序列化
	if err := c.BodyParser(&model); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON body")
	}

	err := Pubweb(model)
	if err != nil {
		return Fail(c, err.Error())
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
		return Fail(c, err.Error())
	}
	return AppResult(c).Success(model,"部署java完成")

}
