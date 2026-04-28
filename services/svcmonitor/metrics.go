package main

import (
	"fmt"
	"net/http"
)

func serveMetrics(w http.ResponseWriter, services map[string]*service) {
	fmt.Fprintf(w, "# HELP UP (1) or DOWN (0)\n")
	fmt.Fprintf(w, "# TYPE service_up gauge\n")
	for _, s := range services {
		s.mutex.Lock()
		fmt.Fprintf(w, "service_up{service=\"%s\"} %d\n\n", s.name, s.status)
		s.mutex.Unlock()
	}
}
