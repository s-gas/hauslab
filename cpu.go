package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
)

func getCPUUsage() {
	percent, _ := cpu.Percent(0, false)
	fmt.Printf("CPU usage:\t%.2f%%\n", percent[0])
}
