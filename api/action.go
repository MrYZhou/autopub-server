package api

import (
	. "log/slog"

	. "github.com/MrYZhou/outil/command"
	. "github.com/MrYZhou/outil/ssh"
)

/**
发布web应用
*/
func Pubweb(localPath string,remotePath string) {
	con :=  Myserver()
	defer con.Client.Close()
	defer con.SftpClient.Close()
	Info("开始打包")
	Run(localPath, "npm run build")
	Info("开始上传")
	con.UploadDir(localPath+"/dist",remotePath)
	Info("上传完毕")
}
/**

*/
func Pubjava(javaProjectPath string, localJarPath string,remotePath string) error{
	con :=  Myserver()
	defer con.Client.Close()
	defer con.SftpClient.Close()
	Info("开始打包")
	// 如果有mvnd使用mvnd
	Run(javaProjectPath, "mvn clean -Dfile.encoding=UTF-8 package")
	Info("开始上传")
	con.UploadFile(localJarPath,remotePath)
	con.Run("echo success")
	Info("上传完毕")
	return nil
}

func TestTransfer(localPath,remotePath string) {
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
