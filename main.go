package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	endpoint := "/metrics"
	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		CPUUsage, err := getCPUUsage()
		if err != nil {
			log.Println("Failed to get CPU usage")
			return
		}
		RAMUsage, err := getRAMUsage()
		if err != nil {
			log.Println("Failed to get RAM usage")
			return
		}
		fmt.Fprintf(w, "cpu_usage_percent %.2f\n", CPUUsage)
		fmt.Fprintf(w, "ram_usage_percent %.2f\n", RAMUsage)
	})

	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("traces listening on %s%s\n", addr, endpoint)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
