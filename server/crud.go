package api

// 此文件主要写各种模型,和crud的操作

// java应用的模型
type JarUpload struct {
	JavaProjectPath string `json:"javaProjectPath"` // java项目根路径
	LocalJarPath    string `json:"localJarPath"`    // 生成的jar文件路径
	RemotePath      string `json:"remotePath"`      // 远程路径
	PubCommand      string `json:"pubCommand"`      // 发布命令或打包命令
	ExecCommand			string `json:"execCommand"`     // 后置发布命令
}

// web应用的模型
type WebrUpload struct {
	LocalPath  string `json:"localPath"`  // 本地web项目路径
	RemotePath string `json:"remotePath"` // 远程web项目路径
}

func create() {
}
