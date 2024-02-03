package router

import (
	. "autopub-server/server"
	. "autopub-server/util"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/gofiber/fiber/v2"
)

// 代码生成相关
func init(){
	app := App()
	// 创建子路由
	api := app.Group("/gen")
	api.Get("/", gen)

	api.Get("/list", genlist)
}
func genlist(c *fiber.Ctx) error {
	config, _ := gplus.SelectList[Config](nil)
  
	return AppResult(c).Success(config)
}

func gen(c *fiber.Ctx) error {
	
	return AppResult(c).Success()
}
