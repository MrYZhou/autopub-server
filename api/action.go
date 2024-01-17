package api

import (
	. "log/slog"

	. "github.com/MrYZhou/outil/command"
	. "github.com/MrYZhou/outil/ssh"
)

/**
发布web应用
*/
func Pubweb(webBase string) {
	
	con :=  Myserver()
	defer con.Client.Close()
	defer con.SftpClient.Close()
	Info("开始打包")
	Run(webBase, "npm run build")
	Info("开始上传")
	con.UploadDir(webBase+"/dist","/root/testweb")
	Info("上传完毕")
}

func TestTransfer(localPath,serverPath string) {
	Info("转移") 
}

var host = "192.168.0.62:22"
var user="root"
var password="YH4WfLbGPasRLVhs"
// 初始化环境
func Myserver() *Cli{
	con  ,_ := Server(host, user, password)
	return con
}
