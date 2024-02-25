package server

//对象模型
type Host struct {
	Id  string  `json:"id"` 
	Name string `json:"name"` // 主机名
	Host string `json:"host"` // 主机ip
	Port string `json:"port"` // 端口
	Password string `json:"password"` // 密码
	SecretValue string `json:"secretValue"` // 密钥
}