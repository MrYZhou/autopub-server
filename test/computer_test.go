package main

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"testing"
	"time"
)

func TestCpu(t *testing.T) {
	for i := 0; i < 1; i++ {
		cpuUsage, err := cpu.Percent(time.Second, false)
		if err != nil {
			fmt.Println("error getting cpu usage:", err)
			return
		}
		fmt.Printf("CPU usage: %f%%\n", cpuUsage[0])
		time.Sleep(time.Second)
	}
}
