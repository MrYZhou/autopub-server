package router

import (
	. "autopub-server/server"
	. "autopub-server/util"
	"net/url"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/gofiber/fiber/v2"
)

// 服务器资源保存
func init() {
	app := App()
	api := app.Group("/host")
	api.Post("/list", hostlist)
	api.Get("/get/:id", hostget)
	api.Post("/detail/:id", hostdetail)
	api.Post("/add", hostadd)
	api.Get("/delete/:id", hostudelete)
	api.Post("/update", hostupdate)
	api.Post("/export", hostexport)
	api.Post("/import", hostimport)
}

func hostlist(c *fiber.Ctx) error {
	host, _ := gplus.SelectList[Host](nil)

	return AppResult(c).Success(host)
}

func hostdetail(c *fiber.Ctx) error {
	host, _ := gplus.SelectList[Host](nil)

	return AppResult(c).Success(host)
}

func hostget(c *fiber.Ctx) error {
	id := c.Params("id")
	host, _ := gplus.SelectById[Host](id)

	return AppResult(c).Success(host)
}
func hostadd(c *fiber.Ctx) error {
	var model Host

	if err := c.BodyParser(&model); err != nil {
		return AppResult(c).Fail("请求体数据解析错误")
	}
	hostValues := url.Values{}
	hostValues.Set("host", model.Host)
	host, _ := gplus.SelectList[Host](
		gplus.BuildQuery[Host](hostValues))
	if len(host) > 0 {
		return AppResult(c).Fail("已经存在该主机")
	}
	model.Id = GetId()
	gplus.Insert[Host](&model)

	return AppResult(c).Success("添加成功")
}
func hostudelete(c *fiber.Ctx) error {
	gplus.DeleteById[Host]("1")

	return AppResult(c).Success("删除成功")
}
func hostupdate(c *fiber.Ctx) error {
	host := Host{
		Name: "test",
		Id:   "1",
	}
	gplus.UpdateById[Host](&host)

	return AppResult(c).Success("更新成功")
}

func hostimport(c *fiber.Ctx) error {
	host, _ := gplus.SelectList[Host](nil)

	return AppResult(c).Success(host)
}
func hostexport(c *fiber.Ctx) error {
	host, _ := gplus.SelectList[Host](nil)

	return AppResult(c).Success(host)
}
