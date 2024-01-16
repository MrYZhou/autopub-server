package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	. "github.com/MrYZhou/outil/command"
	"github.com/gofiber/fiber/v2"

	// "github.com/MrYZhou/outil/file"
	// "github.com/MrYZhou/outil/ssh"
	// 用点的意思是可以不用模块在点方法,否则要file.某个方法 ssh.某个方法
	. "autopub-server/file"
	. "autopub-server/ssh"
)

// 初始化环境
func InitEnv() {
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

func main() {
	// 创建一个新的 Fiber 应用实例
	app := fiber.New()
	// 初始化
	InitEnv()
	app.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("autopub server")
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

	app.Get("/help", func(c *fiber.Ctx) error {

		return c.SendString(`
		设计思路:
		1.先在服务器安装docker,nginx
		2.在服务器编写nginx配置
		3.前置条件环境资源（一般就是服务器连接配置）
		4.把每一项作为一个env item(环境项).新建一个任务，任务包含多个环境项构成的环境链,
		会依次执行.新建一个环境项需要选择环境信息。
		`)
	})

	// 设置服务器监听地址和端口
	if err := app.Listen("127.0.0.1:8083"); err != nil {
		// 如果监听失败，则输出错误信息并终止程序
		panic(err)
	}
}

func  Help(command string) {
	if strings.Contains(command, "help") == true {
		fmt.Println(`使用步骤:
		1.先在服务器安装docker,nginx
		2.在服务器编写nginx配置
		3.新建.env配置文件, 然后 o config .
		4.o pub web 发布前端, o pub java 发布后端`)
	}
}
func  Info(command string) {
	if strings.Contains(command, "info") == true {
		fmt.Println("create by larry!")
	}
}
func  Pub(pubType string) {
	c, _ := Server(os.Getenv("host"), os.Getenv("user"), os.Getenv("password"))

	defer c.Client.Close()
	defer c.SftpClient.Close()

	if pubType == "all" {
		PubType("java", c)
		PubType("web", c)
	} else {
		PubType(pubType, c)
	}
}
func PubType(name string, c *Cli) {
	packageCode(name)
	pubCode(name, c)
}

/*
上传文件到服务器部署

pubType 部署的类型 web,java,all
*/
func pubCode(pubType string, c *Cli) {

	if pubType == "web" {
		fmt.Println("上传前端文件")
		base := path.Join(os.Getenv("webBase"), "dist")
		target := os.Getenv("webTarget")
		c.UploadDir(base, target)
	} else if pubType == "java" {
		fmt.Println("上传jar文件")
		jarFilePath := os.Getenv("jarFilePath")

		// 获取jarFilePath的jar文件名
		file, _ := os.Open(jarFilePath)
		name := file.Name()

		remoteJarHome := os.Getenv("remoteJarHome")
		remoteJarFilePath := remoteJarHome + name
		fileList := c.SliceUpload(remoteJarHome, jarFilePath, 6)
		c.ConcatRemoteFile(fileList, remoteJarFilePath)
		c.Run("rm -rf " + strings.Join(fileList, " "))
		// 镜像构建
		init := InitDockerfile(c, remoteJarHome, name)

		// 运行容器
		RunContainer(init, c)

	}
	fmt.Println("部署完成")
}

// 保存配置
func  Config(filePath string) {
	if filePath == "." {
		filePath = "./.env"
	}
	TransFile(filePath, "C:/oenv/.env")
	fmt.Println("配置完成")
}

/*
init 没有生成过dockerfile文件,init为false
*/
func RunContainer(init bool, c *Cli) {
	fmt.Println("运行容器")
	direct := ""
	javaContainerName := os.Getenv("javaContainerName")
	imageName := os.Getenv("imageName")
	remoteJarHome := os.Getenv("remoteJarHome")
	port := os.Getenv("port") + ":" + os.Getenv("port")

	if init == false {
		// 不需要输出,下面两行考虑到容器名可能已经存在,需要先移除
		c.RunQuiet("docker stop " + javaContainerName)
		c.RunQuiet("docker rm " + javaContainerName)
		// 需要映射目录这样restart才有意义
		direct = "docker run -d --name " + javaContainerName + " -p " + port + " -v " + remoteJarHome + ":/java " + imageName
	} else {
		direct = "docker restart " + javaContainerName
	}
	c.Run(direct)
}

/*
remoteJarHome  服务器jar文件所在目录

name jar文件的名字
*/
func InitDockerfile(c *Cli, remoteJarHome string, name string) bool {
	dockerFilePath := path.Join(remoteJarHome, "Dockerfile")
	init := c.IsFileExist(dockerFilePath)
	if init == false {
		// 创建dockerfile文件
		ftpFile, _ := c.CreateFile(dockerFilePath)

		version := os.Getenv("jdk")
		port := os.Getenv("port")
		if version == "" {
			version = "8"
		}

		b := []byte("FROM openjdk:" + version + "-slim" + "\n")
		ftpFile.Write(b)
		b = []byte("WORKDIR /java" + "\n")
		ftpFile.Write(b)
		// 因为使用-v映射方式,不需要直接添加进去
		// b = []byte("ADD *.jar /java/app.jar" + "\n")
		// ftpFile.Write(b)
		b = []byte(`ENTRYPOINT ["java","-jar","/java/` + name + `"]` + "\n")
		ftpFile.Write(b)
		b = []byte("EXPOSE " + port)
		ftpFile.Write(b)
		imageName := os.Getenv("imageName")
		fmt.Println("正在构建镜像")
		build := "docker build -f " + dockerFilePath + " -t  " + imageName + " " + remoteJarHome
		msg, err := c.Run(build)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(msg)
		fmt.Println("构建完成")

	}
	return init
}
