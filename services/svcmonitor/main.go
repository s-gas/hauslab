package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	services := getServices()

	var mutex sync.Mutex

	endpoint := "/metrics"
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		serveMetrics(w, &mutex, services)
	})

	go checkServices(services, &mutex)

	port := 1024
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("svcmonitor listening on %s%s\n", addr, endpoint)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
