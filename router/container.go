package router

import (
	. "autopub-server/server"
	. "autopub-server/util"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/gofiber/fiber/v2"
)

// 容器操作
func init(){
	app := App()
	api := app.Group("/container")
	api.Post("/list", containerlist)
	api.Get("/get/:id",containerget)
	api.Post("/detail/:id", containerdetail)
	api.Post("/add", containeradd)
	api.Get("/delete/:id", containerudelete)
	api.Post("/update", containerupdate)
	api.Post("/export", containerexport)
	api.Post("/import", containerimport)
}

func containerlist(c *fiber.Ctx) error {
	container, _ := gplus.SelectList[Container](nil)
  
	return AppResult(c).Success(container)
}

func containerdetail(c *fiber.Ctx) error {
	container, _ := gplus.SelectList[Container](nil)
  
	return AppResult(c).Success(container)
}

func containerget(c *fiber.Ctx) error {
	id:=c.Params("id")
	container, _ := gplus.SelectById[Container](id)
  
	return AppResult(c).Success(container)
}
func containeradd(c *fiber.Ctx) error {
	var model Container
	if err := c.BodyParser(&model); err != nil {
		return AppResult(c).Fail("Invalid JSON body")
	}
	gplus.Insert[Container](&model)
  
	return AppResult(c).Success("添加成功")
}
func containerudelete(c *fiber.Ctx) error {
	gplus.DeleteById[Container]("1")
  
	return AppResult(c).Success("删除成功")
}
func containerupdate(c *fiber.Ctx) error {
	container := Container{
		Name: "test",
		Id: "1",
	}
	gplus.UpdateById[Container](&container)
  
	return AppResult(c).Success("更新成功")
}

func containerimport(c *fiber.Ctx) error {
	container, _ := gplus.SelectList[Container](nil)
  
	return AppResult(c).Success(container)
}
func containerexport(c *fiber.Ctx) error {
	container, _ := gplus.SelectList[Container](nil)
  
	return AppResult(c).Success(container)
}