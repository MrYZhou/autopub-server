package router

import (
	. "autopub-server/server"
	. "autopub-server/util"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/gofiber/fiber/v2"
)
// 服务器资源保存
func init(){
	app := App()
	api := app.Group("/project")
	api.Post("/list", projectlist)
	api.Get("/get/:id",projectget)
	api.Post("/detail/:id", projectdetail)
	api.Post("/add", projectadd)
	api.Get("/delete/:id", projectudelete)
	api.Post("/update", projectupdate)
	api.Post("/export", projectexport)
	api.Post("/import", projectimport)
}

func projectlist(c *fiber.Ctx) error {
	config, _ := gplus.SelectList[Config](nil)
  
	return AppResult(c).Success(config)
}

func projectdetail(c *fiber.Ctx) error {
	config, _ := gplus.SelectList[Config](nil)
  
	return AppResult(c).Success(config)
}

func projectget(c *fiber.Ctx) error {
	id:=c.Params("id")
	config, _ := gplus.SelectById[Config](id)
  
	return AppResult(c).Success(config)
}
func projectadd(c *fiber.Ctx) error {
	var model project
	if err := c.BodyParser(&model); err != nil {
		return AppResult(c).Fail("Invalid JSON body")
	}
	gplus.Insert[project](&model)
  
	return AppResult(c).Success("添加成功")
}
func projectudelete(c *fiber.Ctx) error {
	gplus.DeleteById[Config]("1")
  
	return AppResult(c).Success("删除成功")
}
func projectupdate(c *fiber.Ctx) error {
	config := Config{
		Name: "test",
		Id: "1",
	}
	gplus.UpdateById[Config](&config)
  
	return AppResult(c).Success("更新成功")
}

func projectimport(c *fiber.Ctx) error {
	config, _ := gplus.SelectList[Config](nil)
  
	return AppResult(c).Success(config)
}
func projectexport(c *fiber.Ctx) error {
	config, _ := gplus.SelectList[Config](nil)
  
	return AppResult(c).Success(config)
}