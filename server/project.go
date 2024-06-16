package server

//对象模型
type Project struct {
	Id       string `json:"id"`
	Name     string `json:"name"`     // 项目名
	Pid      string `json:"host"`     // 项目归属id
	Document string `json:"document"` // 项目资料
}
