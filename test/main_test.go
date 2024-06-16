package main

import (
	. "log/slog"
	"testing"

	. "github.com/MrYZhou/outil/command"
	. "github.com/MrYZhou/outil/ssh"
)

func TestTransfer(t *testing.T) {
	Info("转移")
}

// 初始化环境
func Myserver() *Cli {
	host := "192.168.0.62:22"
	user := "root"
	password := "YH4WfLbGPasRLVhs"
	con, _ := Server(host, user, password)
	return con
}

/**
发布web应用
*/
func Pubweb(webBase string) {

	con := Myserver()
	defer con.Client.Close()
	defer con.SftpClient.Close()

	Run(webBase, "npm run build")
}
