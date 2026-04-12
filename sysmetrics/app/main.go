package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	endpoint := "/metrics"
	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		cpu, err := getCPUUsage()
		if err != nil {
			log.Println("Failed to get CPU usage")
			return
		}
		ram, err := getRAMUsage()
		if err != nil {
			log.Println("Failed to get RAM usage")
			return
		}
		writeMetrics(w, cpu, ram)
	})

	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("sysmetrics listening on %s%s\n", addr, endpoint)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
