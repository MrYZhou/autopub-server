package server

//对象模型
type Task struct {
	Id      string `json:"id"`
	Name    string `json:"name"`    // 任务名
	Type    string `json:"type"`    // 任务类型
	Content string `json:"content"` // 任务内容
}
