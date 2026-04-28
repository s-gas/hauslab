package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	services := getServices()

	endpoint := "/metrics"
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		serveMetrics(w, services)
	})

	go checkServices(services)

	port := 1025
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("svcmonitor listening on %s%s\n", addr, endpoint)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
