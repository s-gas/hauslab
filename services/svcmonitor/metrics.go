package main

import (
	"fmt"
	"net/http"
	"sync"
)

func serveMetrics(w http.ResponseWriter, mutex *sync.Mutex, services map[string]service) {
	mutex.Lock()
	fmt.Fprintf(w, "# HELP UP (1) or DOWN (0)\n")
	fmt.Fprintf(w, "# TYPE service_up gauge\n")
	for _, s := range services {
		fmt.Fprintf(w, "service_up{service=\"%s\"} %d\n\n", s.name, s.status)
	}
	mutex.Unlock()
}
