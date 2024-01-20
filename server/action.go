package api

import (
	. "log/slog"

	. "github.com/MrYZhou/outil/command"
	. "github.com/MrYZhou/outil/ssh"
)

// 此文件主要写各种事件的动作

/*
发布web应用
*/
func Pubweb(model WebrUpload) error {
	localPath := model.LocalPath
	remotePath := model.RemotePath
	con := Myserver()
	defer con.Client.Close()
	defer con.SftpClient.Close()
	Info("开始打包")
	err := Run(localPath, "npm run build")
	if err != nil {
		return err
	}
	Info("开始上传")
	con.UploadDir(localPath+"/dist", remotePath)
	Info("上传完毕")
	return nil
}

/*
*
发布java应用
*/
func Pubjava(model JarUpload) error {
	javaProjectPath := model.JavaProjectPath
	localJarPath := model.LocalJarPath
	remotePath := model.RemotePath
	pubCommand := model.PubCommand
	if pubCommand == "" {
		pubCommand = "mvn -Dfile.encoding=UTF-8 package"
	}
	con := Myserver()
	defer con.Client.Close()
	defer con.SftpClient.Close()
	Info("开始打包")
	Run(javaProjectPath, pubCommand)
	Info("开始上传")
	con.UploadFile(localJarPath, remotePath)
	con.Run("cd /www/wwwroot/crm2.test.project.jnpf.work/java/jnpf-admin && ./server.sh restart")
	Info("上传完毕")
	return nil
}

var host = "192.168.0.62:22"
var user = "root"
var password = "YH4WfLbGPasRLVhs"

// 初始化环境
func Myserver() *Cli {
	con, _ := Server(host, user, password)
	return con
}
