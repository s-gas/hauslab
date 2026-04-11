package main

import (
	"github.com/shirou/gopsutil/v3/mem"
)

func getRAMUsage() (float64, error) {
	ram, err := mem.VirtualMemory()
	if err != nil {
		return 0.0, err
	}
	return ram.UsedPercent, nil
}
