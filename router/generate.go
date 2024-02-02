package router

import (
	"github.com/gofiber/fiber/v2"

	. "autopub-server/server"
	. "autopub-server/util"
)
func Router(){
}
// 代码生成相关
func init(){
	app := App()
	// 创建子路由
	api := app.Group("/gen")
	{
		api.Get("/a", func(c *fiber.Ctx) error {
			println(c)
			return nil
		})
	}
}

func Generate(c *fiber.Ctx) error {
	
	return AppResult(c).Success()
}
