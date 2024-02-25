package server

//对象模型
type Container struct {
	ContainerId  string  `json:"containerId"` // 容器id  
	Name string `json:"name"` // 容器名
}