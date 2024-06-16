package server

//对象模型
type Host struct {
	Id          string `json:"id"`
	Name        string `json:"name"`        // 主机别名
	Host        string `json:"host"`        // 主机ip
	Port        string `json:"port"`        // 端口
	Account     string `json:"account"`     // 登陆用户
	Password    string `json:"password"`    // 密码
	SecretValue string `json:"secretValue"` // 密钥
}
