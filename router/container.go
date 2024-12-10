package router

import (
	. "autopub-server/util"
	"context"
	"log"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/gofiber/fiber/v2"
)

// 容器管理
func init() {
	app := App()
	api := app.Group("/container")
	api.Post("/list", containerlist)
	api.Get("/get/:id", containerget)
	api.Post("/detail/:id", containerdetail)
	api.Post("/add", containeradd)
	api.Get("/delete/:id", containerudelete)
	api.Post("/update", containerupdate)
	api.Post("/export", containerexport)
	api.Post("/import", containerimport)
}

// 创建一个Docker客户端实例
func getCli() (*client.Client, error) {

	// docker20.x版本匹配的sdk是旧版的1.41
	// opts := []client.Opt{
	// 	client.WithVersion("1.41"),
	// }
	// cli, err := client.NewClientWithOpts(opts...)

	// 新版本连接写法
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Println("Error creating docker client:", err)
	}
	return cli, err
}

// 对象模型
type Container struct {
	ContainerId string `json:"containerId"` // 容器id
	Name        string `json:"name"`        // 容器名
	State       string `json:"state"`       // 运行状态
}

func containerlist(c *fiber.Ctx) error {
	// todo 获取不同主机的信息,遍历主机获取所有docker运行容器
	// container, _ := gplus.SelectList[Container](nil)
	cli, err := getCli()
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
			State:       container.State,
		}
		mycontainer = append(mycontainer, con)
	}
	return AppResult(c).Success(mycontainer)
}

func containerdetail(c *fiber.Ctx) error {
	container, _ := gplus.SelectList[Container](nil)

	return AppResult(c).Success(container)
}
func containerstart(c *fiber.Ctx) error {
	// 镜像id
	id := c.Params("id")
	cli, _ := getCli()
	cli.ContainerStart(context.Background(), id, container.StartOptions{})
	return AppResult(c).Success()
}
func containerrestart(c *fiber.Ctx) error {
	// 容器id
	id := c.Params("id")
	cli, _ := getCli()
	cli.ContainerRestart(context.Background(), id, container.StopOptions{})
	return AppResult(c).Success()
}

func containerget(c *fiber.Ctx) error {
	id := c.Params("id")
	container, _ := gplus.SelectById[Container](id)

	return AppResult(c).Success(container)
}
func containeradd(c *fiber.Ctx) error {
	var model Container
	if err := c.BodyParser(&model); err != nil {
		return AppResult(c).Fail("请求体数据解析错误")
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
