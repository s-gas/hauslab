package main

import (
	"fmt"
	"net/http"
)

func serveMetrics(w http.ResponseWriter, services map[string]*Service) {
	fmt.Fprintf(w, "# HELP service_up Status of the service (1 for UP, 0 for DOWN)\n")
	fmt.Fprintf(w, "# TYPE service_up gauge\n")
	for name, s := range services {
		s.mutex.Lock()
		fmt.Fprintf(w, "service_up{service=\"%s\"} %d\n", name, s.status)
		s.mutex.Unlock()
	}
}
