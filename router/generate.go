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

	api.Post("/list", genlist)
	api.Get("get",genget)
	api.Post("/detail", gendetail)

	api.Post("/add", genadd)
	api.Post("/delete", genudelete)
	api.Post("/update", genupdate)

	api.Post("/export", genexport)
	api.Post("/import", genimport)

}

func genimport(c *fiber.Ctx) error {
	config, _ := gplus.SelectList[Config](nil)
  
	return AppResult(c).Success(config)
}
func genexport(c *fiber.Ctx) error {
	config, _ := gplus.SelectList[Config](nil)
  
	return AppResult(c).Success(config)
}
func genget(c *fiber.Ctx) error {
	config, _ := gplus.SelectList[Config](nil)
  
	return AppResult(c).Success(config)
}
func genadd(c *fiber.Ctx) error {
	config, _ := gplus.SelectList[Config](nil)
  
	return AppResult(c).Success(config)
}
func genudelete(c *fiber.Ctx) error {
	config, _ := gplus.SelectList[Config](nil)
  
	return AppResult(c).Success(config)
}
func genupdate(c *fiber.Ctx) error {
	config, _ := gplus.SelectList[Config](nil)
  
	return AppResult(c).Success(config)
}
func genlist(c *fiber.Ctx) error {
	config, _ := gplus.SelectList[Config](nil)
  
	return AppResult(c).Success(config)
}

func gendetail(c *fiber.Ctx) error {
	config, _ := gplus.SelectList[Config](nil)
  
	return AppResult(c).Success(config)
}
func gen(c *fiber.Ctx) error {
	
	return AppResult(c).Success()
}
