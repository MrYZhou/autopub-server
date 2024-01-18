package api

type JarUpload struct {
	JavaProjectPath string `json:"javaProjectPath"` // java项目根路径
	LocalJarPath    string `json:"localJarPath"`    // 生成的jar文件路径
	RemotePath      string `json:"remotePath"`      // 远程路径
	PubCommand      string `json:"pubCommand"`      // 发布命令或打包命令
	Msg             string `json:"msg"`             // 提示信息
}

type WebrUpload struct {
	LocalPath  string `json:"localPath"`  // 本地web项目路径
	RemotePath string `json:"remotePath"` // 远程web项目路径
	Msg        string `json:"msg"`        // 提示信息
}

func create() {
}
