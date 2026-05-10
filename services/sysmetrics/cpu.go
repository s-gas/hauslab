package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

func getCpuUsage() (float64, error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0.0, fmt.Errorf("getCpuUsage: %w", err)
	}
	return percent[0], nil
}
