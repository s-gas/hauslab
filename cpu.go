package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
)

func getCPU() {
	percent, _ := cpu.Percent(0, false)
	fmt.Printf("CPU usage: %.2f\n", percent[0])
}
