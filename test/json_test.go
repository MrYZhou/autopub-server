package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Blog struct {
	BlogId  string `json:"blogId"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Uid     int32  `json:"uid"`
	State   int32  `json:"state"`
}

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"` // 使用omitempty选项表示当字段为空时不在JSON中输出
}

func TestJsonDecode(t *testing.T) {
	str := `{"title": "some title","Uid":1}`
	var data Blog
	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		t.Errorf("Error during JSON unmarshalling: %v", err)
		return
	}
	fmt.Println(data)
}

func TestJsonEncode(t *testing.T) {
	// 创建一个Person实例
	p := Person{
		Name:  "张三",
		Age:   25,
		Email: "",
	}
	jsonBytes, _ := json.Marshal(p)
	jsonString := string(jsonBytes)
	fmt.Println(jsonString)
}
