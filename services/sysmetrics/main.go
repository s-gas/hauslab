package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	endpoint := "/metrics"
	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		cpu, err := getCpuUsage()
		if err != nil {
			log.Println(err)
			return
		}
		ram, err := getRamUsage()
		if err != nil {
			log.Println(err)
			return
		}
		writeMetrics(w, cpu, ram)
	})

	port := 1024
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("metrics-server listening on %s%s\n", addr, endpoint)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
