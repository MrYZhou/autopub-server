package router

import (
	. "autopub-server/server"
	. "autopub-server/util"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/gofiber/fiber/v2"
)

// 服务器资源保存
func init() {
	app := App()
	api := app.Group("/project")
	api.Post("/list", projectlist)
	api.Get("/get/:id", projectget)
	api.Post("/detail/:id", projectdetail)
	api.Post("/add", projectadd)
	api.Get("/delete/:id", projectudelete)
	api.Post("/update", projectupdate)
	api.Post("/export", projectexport)
	api.Post("/import", projectimport)
}

func projectlist(c *fiber.Ctx) error {
	project, _ := gplus.SelectList[Project](nil)

	return AppResult(c).Success(project)
}

func projectdetail(c *fiber.Ctx) error {
	project, _ := gplus.SelectList[Project](nil)

	return AppResult(c).Success(project)
}

func projectget(c *fiber.Ctx) error {
	id := c.Params("id")
	project, _ := gplus.SelectById[Project](id)

	return AppResult(c).Success(project)
}
func projectadd(c *fiber.Ctx) error {
	var model Project
	if err := c.BodyParser(&model); err != nil {
		return AppResult(c).Fail("请求体数据解析错误")
	}
	gplus.Insert[Project](&model)

	return AppResult(c).Success("添加成功")
}
func projectudelete(c *fiber.Ctx) error {
	gplus.DeleteById[Project]("1")

	return AppResult(c).Success("删除成功")
}
func projectupdate(c *fiber.Ctx) error {
	project := Project{
		Name: "test",
		Id:   "1",
	}
	gplus.UpdateById[Project](&project)

	return AppResult(c).Success("更新成功")
}

func projectimport(c *fiber.Ctx) error {
	project, _ := gplus.SelectList[Project](nil)

	return AppResult(c).Success(project)
}
func projectexport(c *fiber.Ctx) error {
	project, _ := gplus.SelectList[Project](nil)

	return AppResult(c).Success(project)
}
