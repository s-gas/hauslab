package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
)

func getRamUsage() (float64, error) {
	ram, err := mem.VirtualMemory()
	if err != nil {
		return 0.0, fmt.Errorf("getRamUsage: %w", err)
	}
	return ram.UsedPercent, nil
}
