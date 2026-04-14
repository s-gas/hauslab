package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	services := getServices()
	logFileName := "svcmonitor.log"
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open %s\n", logFileName)
		os.Exit(1)
	}
	for {
		for _, s := range services {
			resp, err := http.Get(s.domain)
			if err != nil {
				fmt.Fprintf(logFile, "%s is down\n", s.name)
			}
			if resp.StatusCode != 200 {
				fmt.Fprintf(logFile, "%s is down\n", s.name)
			} else {
				fmt.Fprintf(logFile, "%s is up\n", s.name)
			}
			time.Sleep(1 * time.Second)
		}
		time.Sleep(5 * time.Minute)
	}
}
