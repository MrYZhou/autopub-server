package router

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"

	. "autopub-server/util"

	. "github.com/MrYZhou/outil/command"
	. "github.com/MrYZhou/outil/ssh"
)

// 部署应用
func init() {
	app := App()

	// 创建子路由
	api := app.Group("/action")
	api.Post("start", start)
}
func start(c fiber.Ctx) error {
	var model ActionModel
	// 从请求体中读取JSON内容并反序列化
	if err := c.Bind().Body(&model); err != nil {
		return AppResult(c).Fail("请求体数据解析错误")
	}
	con := Getserver()
	defer con.Client.Close()
	defer con.SftpClient.Close()
	for _, action := range model.Actions {
		log.Println("\n<========== " + action.Name + "开始 ==========>")
		err := DoAction(con, action)
		if err != nil {
			log.Println("\n<========== " + action.Name + "执行失败 ==========>")
			break
		} else {
			log.Println("\n<========== " + action.Name + "结束 ==========>")
		}
	}
	return AppResult(c).Success()
}
func DoAction(cli *Cli, action Action) error {
	var actionType = action.ActionType
	if actionType == "uploadDir" {
		cli.UploadDir(action.LocalPath, action.RemotePath)
	} else if actionType == "uploadFile" {
		cli.UploadFile(action.LocalPath, action.RemotePath)
	} else {
		switch execCommand := action.ExecCommand.(type) {
		case string:
			err := Run(action.WorkPath, execCommand)
			if err != nil {
				return err
			}
		case []interface{}:
			for _, value := range execCommand {
				err := Run(action.WorkPath, value.(string))
				if err != nil {
					return err
				}
			}
		}
	}

	switch postCommand := interface{}(action.PostCommand).(type) {
	case string:
		log.Println("\n<========== 执行后置命令:" + postCommand + "==========>")
		cli.Run(postCommand)
	case []interface{}:
		for _, value := range postCommand {
			log.Println("\n<========== 执行后置命令:" + value.(string) + "==========>")
			cli.Run(value.(string))
		}
	}
	return nil
}

type ActionModel struct {
	Actions []Action `json:"actions"` // 操作步骤
}

type Action struct {
	Name        string      `json:"name"`        // 步骤名称
	WorkPath    string      `json:"workPath"`    // 工作路径,此项不空则会在此目录为根执行操作
	LocalPath   string      `json:"localPath"`   // 本地路径
	RemotePath  string      `json:"remotePath"`  // 远程路径
	ActionType  string      `json:"actionType"`  // 动作类型, upload,common, 默认是common
	ExecCommand interface{} `json:"execCommand"` // 命令组,需要按顺序执行的命令
	ServerId    string      `json:"serverId"`    // 服务器id,用来标识用哪一个服务器来执行动作,保留key
	PostCommand interface{} `json:"postCommand"` // 服务器上传后执行什么命令
}

// 初始化环境
func Getserver() *Cli {
	accessType := os.Getenv("accessType")
	if accessType == "password" {
		con, _ := Server(os.Getenv("host"), os.Getenv("user"), os.Getenv("password"))
		return con
	} else {
		contentBytes, err := os.ReadFile(os.Getenv("rsaFilePath"))
		if err != nil {
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
