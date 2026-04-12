package main

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

func getCPUUsage() (float64, error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0.0, err
	}
	return percent[0], nil
}
