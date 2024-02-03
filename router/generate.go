package router

import (
	"github.com/gofiber/fiber/v2"

	. "autopub-server/util"
)

// 代码生成相关
func init(){
	app := App()
	// 创建子路由
	api := app.Group("/gen")
	api.Get("/", gen)
}

func gen(c *fiber.Ctx) error {
	
	return AppResult(c).Success()
}
