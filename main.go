package main

import (
	"fmt"
	"os"

	. "github.com/MrYZhou/outil/command"
	"github.com/gofiber/fiber/v2"

	. "github.com/MrYZhou/outil/ssh"

	. "autopub-server/api"
)

var con Cli

// 初始化环境
func InitEnv() {
	// docker 初始化.
	// con  ,_ := Server(os.Getenv("host"), os.Getenv("user"), os.Getenv("password"))
	// fmt.Println(con)
}

// 执行打包命令
func packageCode(pubType string) {

	fmt.Println("开始打包")
	if pubType == "web" {
		Run(os.Getenv("webBase"), "npm run build")
	} else if pubType == "java" {
		Run(os.Getenv("javaProjectPath"), "mvn clean -Dmaven.test.skip=true package")
	}
}

// model实体定义
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type JarUpload struct {
	JavaProjectPath string  `json:"javaProjectPath"`
	LocalJarPath string `json:"localJarPath"`
	RemotePath string `json:"remotePath"`
}

func main() {
	// 创建一个新的 Fiber 应用实例
	app := fiber.New()
	// 初始化
	InitEnv()
	app.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("autopub server")
	})
	app.Get("pubweb",func(c *fiber.Ctx) error {
		Pubweb("D:/Users/JNPF/Desktop/jnpf-crm2-web","/root/testweb")
		return c.SendString("ok")
	})

	app.Post("pubjava",func(c *fiber.Ctx) error {
		var model JarUpload
		// 从请求体中读取JSON内容并反序列化
		if err := c.BodyParser(&model); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON body")
		}
		Pubjava(model.JavaProjectPath,model.LocalJarPath,model.RemotePath)
		return c.SendString("ok")
	})
	// 创建一个处理POST JSON请求的路由
	app.Post("/user", func(c *fiber.Ctx) error {
		var user User // 创建一个新的User实例用于接收解析后的JSON数据

		// 从请求体中读取JSON内容并反序列化
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON body")
		}

		// 返回响应
		return c.JSON(user)
	})

	// 设置服务器监听地址和端口
	if err := app.Listen("127.0.0.1:8083"); err != nil {
		// 如果监听失败，则输出错误信息并终止程序
		panic(err)
	}
}
