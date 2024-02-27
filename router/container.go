package router

import (
	. "autopub-server/server"
	. "autopub-server/util"
	"context"
	"log"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/gofiber/fiber/v2"
)

// 容器操作
func init() {
	app := App()
	api := app.Group("/container")
	api.Post("/list", containerlist)
	api.Get("tes", tes)
	api.Get("/get/:id", containerget)
	api.Post("/detail/:id", containerdetail)
	api.Post("/add", containeradd)
	api.Get("/delete/:id", containerudelete)
	api.Post("/update", containerupdate)
	api.Post("/export", containerexport)
	api.Post("/import", containerimport)
}
func tes(c *fiber.Ctx) error {
	// 创建一个Docker客户端实例
	// docker20版本
	// opts := []client.Opt{
	// 	client.WithVersion("1.41"),
	// }
	// cli, err := client.NewClientWithOpts(opts...)

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Println("Error creating docker client:", err)

	}

	// 调用API获取所有正在运行的容器
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		log.Println("Error listing containers:", err)
	}
	mycontainer := []Container{}

	// 打印所有运行中的容器ID和名称
	for _, container := range containers {
		con := Container{
			ContainerId: container.ID,
			Name:        container.Names[0],
			State: container.State,
		}
		mycontainer = append(mycontainer, con)
	}
	return AppResult(c).Success(mycontainer)
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
	id := c.Params("id")
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
