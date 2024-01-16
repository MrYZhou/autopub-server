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

func TestJson(t *testing.T) {
	str := `{"title": "some title","Uid":1}`
	var data Blog
	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		t.Errorf("Error during JSON unmarshalling: %v", err)
		return
	}
	fmt.Println(data) 
}