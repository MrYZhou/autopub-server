package main

import (
	"fmt"
	. "o/ssh"
	"strings"
	"testing"
	"time"
)

// 测试远程主机的上传
func Test(t *testing.T) {
	bT := time.Now()
	c,_ := Server("121.5.68.243:22", "root", "!123qweA")
	fileName := "/root/goenv/o/main.go"
	fileList := c.SliceUpload("/root/java31", fileName, 5)
	fmt.Println(fileList)
	c.ConcatRemoteFile(fileList, "/root/java31/main.go")
	c.Close()
	eT := time.Since(bT) // 从开始到当前所消耗的时间
	fmt.Println("总耗时:", eT)
}

// 运行docker容器
func TestDockerRun(t *testing.T){
	c,_ := Server("192.168.0.62:22", "root", "JNPF@116898")
	
	c.Run("docker stop javacrmCon")
	c.Run("docker rm javacrmCon")
	cs,err:=c.Run("docker run -d --name javacrmCon -p 30122:30122 javacrm")
	fmt.Println(cs,err)
}

// 测试连通性
func TestConnect(t *testing.T){
	
	c,err:= Server("121.5.68.242:22", "root", "!123qweA")
	if err != nil {
		fmt.Println("连接失败")
	}else{
		fmt.Println(c)
	}
}
// 检测是否有docker环境
func TestHasDocker(t *testing.T){
	c,_ := Server("121.5.68.243:22", "root", "!123qweA")
	cs,_:=c.Run("docker ps")
	if strings.Contains(cs,"command not found")  {
		fmt.Println("不存在")
	
	}else{
		fmt.Println("存在docker环境")
	}
}
// 检查端口占用
func TestPortUse(t *testing.T){
	c,_ := Server("121.5.68.243:22", "root", "!123qweA")
	cs,_:=c.Run("lsof -i:80")
	if len(cs) > 0 {
		fmt.Println("存在")
	
	}
}

