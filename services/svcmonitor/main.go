package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	services := getServices()
	logFileName := "svcmonitor.log"
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open %s\n", logFileName)
		os.Exit(1)
	}
	for {
		checkServices(services, logFile)
		time.Sleep(1 * time.Minute)
	}
}
