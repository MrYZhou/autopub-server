package api

type JarUpload struct {
	JavaProjectPath string `json:"javaProjectPath"`
	LocalJarPath    string `json:"localJarPath"`
	RemotePath      string `json:"remotePath"`
	PubCommand      string `json:"pubCommand"`
	Msg             string `json:"msg"`
}

type WebrUpload struct {
	LocalPath  string `json:"localPath"`
	RemotePath string `json:"remotePath"`
	Msg        string `json:"msg"`
}

func create() {
}
