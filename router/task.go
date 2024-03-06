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
	api := app.Group("/task")
	api.Post("/list", tasklist)
	api.Get("/get/:id",taskget)
	api.Post("/detail/:id", taskdetail)
	api.Post("/add", taskadd)
	api.Get("/delete/:id", taskudelete)
	api.Post("/update", taskupdate)
	api.Post("/export", taskexport)
	api.Post("/import", taskimport)
}

func tasklist(c *fiber.Ctx) error {
	task, _ := gplus.SelectList[Task](nil)
  
	return AppResult(c).Success(task)
}

func taskdetail(c *fiber.Ctx) error {
	task, _ := gplus.SelectList[Task](nil)
  
	return AppResult(c).Success(task)
}

func taskget(c *fiber.Ctx) error {
	id:=c.Params("id")
	task, _ := gplus.SelectById[Task](id)
  
	return AppResult(c).Success(task)
}
func taskadd(c *fiber.Ctx) error {
	task := Task{
		Name: "test",
		Id: "1",
	}
	 gplus.Insert[Task](&task)
  
	return AppResult(c).Success("添加成功")
}
func taskudelete(c *fiber.Ctx) error {
	gplus.DeleteById[Task]("1")
  
	return AppResult(c).Success("删除成功")
}
func taskupdate(c *fiber.Ctx) error {
	task := Task{
		Name: "test",
		Id: "1",
	}
	gplus.UpdateById[Task](&task)
  
	return AppResult(c).Success("更新成功")
}

func taskimport(c *fiber.Ctx) error {
	task, _ := gplus.SelectList[Task](nil)
  
	return AppResult(c).Success(task)
}
func taskexport(c *fiber.Ctx) error {
	task, _ := gplus.SelectList[Task](nil)
  
	return AppResult(c).Success(task)
}