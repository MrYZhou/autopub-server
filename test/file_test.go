package main

import (
	"testing"

	. "github.com/MrYZhou/outil/file"
)

// 测试本地文件合并
func TestCombineFile(t *testing.T) {
	fileList := SliceFile("/root/goenv/o/test","/root/goenv/o/o.exe",2)
	CombineFile(fileList,"/root/goenv/o/test/o.exe")
}
// 测试本地slice
func TestSliceFile(t *testing.T) {
	SliceFile("/root/goenv/o/test","/root/goenv/o/o.exe",2)
}
