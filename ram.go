package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
)

func getRAMUsage() {
	ram, _ := mem.VirtualMemory()
	fmt.Printf("RAM usage:\t%.2f%%\n", ram.UsedPercent)
}
