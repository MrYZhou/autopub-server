package router

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"

	. "autopub-server/util"

	. "github.com/MrYZhou/outil/ssh"
)

// 部署应用
func init() {
	app := App()

	// 创建子路由
	api := app.Group("/action")
	api.Get("/", action)
	app.Post("start", start)
}
func action(c fiber.Ctx) error {
	return AppResult(c).Success()
}
func start(c fiber.Ctx) error {
	var model ActionModel
	// 从请求体中读取JSON内容并反序列化
	if err := c.Bind().Body(&model); err != nil {
		return AppResult(c).Fail("请求体数据解析错误")
	}
	actions := model.Actions
	for index:=0; index < len(actions) ; index++ {
		action := actions[index]
		fmt.Println(action.WorkPath)
	}
	// cli := Getserver()
	// println(cli)
	return AppResult(c).Success(actions)
}

// java应用的模型
type ActionModel struct {
	Actions []Action `json:"actions"` // 操作步骤
}

type Action struct {
	WorkPath    string   `json:"workPath"`    // 工作路径,此项不空则会在此目录为根执行操作
	LocalPath   string   `json:"localPath"`   // 本地路径
	RemotePath  string   `json:"remotePath"`  // 远程路径
	ActionType  string   `json:"actionType"`  // 动作类型, upload,common, 默认是common
	ExecCommand []string `json:"execCommand"` // 命令组,需要按顺序执行的命令
	ServerId    string   `json:"serverId"`    // 服务器id,用来标识用哪一个服务器来执行动作,保留key
}

// 初始化环境
func Getserver() *Cli {
	accessType := os.Getenv("accessType")
	if accessType == "password" {
		con, _ := Server(os.Getenv("host"), os.Getenv("user"), os.Getenv("password"))
		return con
	} else {
		contentBytes, err := os.ReadFile(os.Getenv("rsaFilePath"))
		if err == nil {
			panic("ssh密钥不存在")
		}
		var cli Cli
		cli.Host = os.Getenv("host")
		cli.User = os.Getenv("user")
		cli.PrivateKey = contentBytes
		con, _ := ConnectServer(cli)
		return con
	}

}
