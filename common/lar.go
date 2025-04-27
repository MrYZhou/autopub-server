package common

import (
	"os/exec"
	"runtime"
	"syscall"
)

func OpenBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{
			HideWindow:    true,       // 关键参数：隐藏子进程窗口
			CreationFlags: 0x08000000, // 可选参数：独立控制台
		}
	}
	return cmd.Start() // 异步执行并返回潜在错误
}
