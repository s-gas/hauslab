package main

import (
	"fmt"
	"os"
)

func main() {
	CPUUsage, err := getCPUUsage()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get CPU usage")
		return
	}
	RAMUsage, err := getRAMUsage()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get RAM usage")
		return
	}
	fmt.Printf("CPU usage:\t%.2f%%\n", CPUUsage)
	fmt.Printf("RAM usage:\t%.2f%%\n", RAMUsage)
}
