package router

import (
	. "autopub-server/util"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/gofiber/fiber/v3"
)

// 代码生成相关
func init() {
	app := App()
	api := app.Group("/gen")
	api.Post("/list", genlist)
	api.Get("/get/:id", genget)
	api.Post("/detail/:id", gendetail)
	api.Post("/add", genadd)
	api.Get("/delete/:id", genudelete)
	api.Post("/update", genupdate)
	api.Post("/export", genexport)
	api.Post("/import", genimport)
}

func genlist(c fiber.Ctx) error {
	config, _ := gplus.SelectList[Config](nil)

	return AppResult(c).Success(config)
}

func gendetail(c fiber.Ctx) error {
	config, _ := gplus.SelectList[Config](nil)

	return AppResult(c).Success(config)
}

func genget(c fiber.Ctx) error {
	id := c.Params("id")
	config, _ := gplus.SelectById[Config](id)

	return AppResult(c).Success(config)
}
func genadd(c fiber.Ctx) error {
	config := Config{
		Name: "test",
		Id:   "1",
	}
	gplus.Insert[Config](&config)

	return AppResult(c).Success("添加成功")
}
func genudelete(c fiber.Ctx) error {
	gplus.DeleteById[Config]("1")

	return AppResult(c).Success("删除成功")
}
func genupdate(c fiber.Ctx) error {
	config := Config{
		Name: "test",
		Id:   "1",
	}
	gplus.UpdateById[Config](&config)

	return AppResult(c).Success("更新成功")
}

func genimport(c fiber.Ctx) error {
	config, _ := gplus.SelectList[Config](nil)

	return AppResult(c).Success(config)
}
func genexport(c fiber.Ctx) error {
	config, _ := gplus.SelectList[Config](nil)

	return AppResult(c).Success(config)
}

// 对象模型
type Config struct {
	Id   string
	Name string
}
