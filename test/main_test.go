package main

import (
	"fmt"
	"testing"

	. "github.com/MrYZhou/outil/command"
	. "github.com/MrYZhou/outil/ssh"
)


func TestTransfer(t *testing.T) {
	fmt.Println("转移") 
}

var con Cli

// 初始化环境
func InitEnv() {
	// docker 初始化.
	host:=""
	user:=""
	password:=""
	con  ,_ := Server(host, user, password)
	defer con.Client.Close()
	defer con.SftpClient.Close()
}
/**
发布web应用
*/
func Pubweb(webBase string) {
	InitEnv()
	fmt.Println(con)
	Run(webBase, "npm run build")
}

func TestPubweb(t *testing.T) {
	Pubweb("")
}