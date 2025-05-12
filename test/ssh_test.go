package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	. "github.com/MrYZhou/outil/ssh"
)

// 测试远程主机的上传
func Test(t *testing.T) {
	bT := time.Now()
	c, _ := Server("121.5.68.243:22", "root", "!123qweA")
	fileName := "/root/goenv/o/main.go"
	fileList := c.SliceUpload("/root/java31", fileName, 5)
	fmt.Println(fileList)
	c.ConcatRemoteFile(fileList, "/root/java31/main.go")
	c.Close()
	eT := time.Since(bT) // 从开始到当前所消耗的时间
	fmt.Println("总耗时:", eT)
}

// 运行docker容器
func TestDockerRun(t *testing.T) {
	c, _ := Server("192.168.0.62:22", "root", "JNPF@116898")

	c.Run("docker stop javacrmCon")
	c.Run("docker rm javacrmCon")
	cs, err := c.Run("docker run -d --name javacrmCon -p 30122:30122 javacrm")
	fmt.Println(cs, err)
}

// 测试连通性密码
func TestConnect(t *testing.T) {
	c, err := Server("192.168.0.68:22", "a06f6d1d64", "4b1d17d5bd")
	if err != nil {
		fmt.Println("连接失败")
	} else {
		fmt.Println(c)
	}
}

// 测试连通性密钥
func TestConnectWithKey(t *testing.T) {
	// 读取私钥文件内容
	contentBytes, _ := os.ReadFile("D:/root/68_rsa")
	var cli Cli
	cli.Host = "192.168.0.68:22"
	cli.User = "root"
	cli.PrivateKey = contentBytes
	con, _ := ConnectServer(cli)
	con.Run("pwd")
}

// 检测是否有docker环境
func TestHasDocker(t *testing.T) {
	c, _ := Server("121.5.68.243:22", "root", "!123qweA")
	cs, _ := c.Run("docker ps")
	if strings.Contains(cs, "command not found") {
		fmt.Println("不存在")

	} else {
		fmt.Println("存在docker环境")
	}
}

// 检查端口占用
func TestPortUse(t *testing.T) {
	c, _ := Server("121.5.68.243:22", "root", "!123qweA")
	cs, _ := c.Run("lsof -i:80")
	if len(cs) > 0 {
		fmt.Println("存在")

	}
}

func TestPubJAVA(t *testing.T) {
	c, _ := Server("121.5.68.243:22", "root", "!123qweA")
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
		direct = "docker run -d --name " + javaContainerName + " -p " + port + " -v " + remoteJarHome + ":/java " + imageName
	} else {
		direct = "docker restart " + javaContainerName
	}
	c.Run(direct)
}
func TestRunContainer(t *testing.T) {
	c, _ := Server("121.5.68.243:22", "root", "!123qweA")
	init := false
	// 运行容器
	RunContainer(init, c)
}
