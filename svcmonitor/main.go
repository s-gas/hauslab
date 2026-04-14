package main

import (
	"fmt"
	"log"
	"net/http"
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
	log.SetOutput(logFile)
	for {
		for _, s := range services {
			resp, err := http.Get(s.domain)
			if err != nil {
				log.Printf("%s is down\n", s.name)
			}
			if resp.StatusCode != 200 {
				log.Printf("%s is down\n", s.name)
			} else {
				log.Printf("%s is up\n", s.name)
			}
			time.Sleep(1 * time.Second)
		}
		time.Sleep(5 * time.Minute)
	}
}
