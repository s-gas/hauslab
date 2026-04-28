package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

type service struct {
	name      string
	domain    string
	statusLog string
	port      int
	status    int
}

func getServices() map[string]service {
	services := make(map[string]service)
	services["sysmetrics"] = service{
		name:      "sysmetrics",
		port:      1024,
		domain:    "http://sysmetrics:1024/metrics",
		status:    0,
		statusLog: "DOWN",
	}
	services["prometheus"] = service{
		name:      "prometheus",
		port:      9090,
		domain:    "http://prometheus:9090",
		status:    0,
		statusLog: "DOWN",
	}
	services["grafana"] = service{
		name:      "grafana",
		port:      3000,
		domain:    "http://grafana:3000",
		status:    0,
		statusLog: "DOWN",
	}
	return services
}

func checkServices(services map[string]service, mutex *sync.Mutex) {
	for {
		mutex.Lock()
		for _, s := range services {
			resp, err := http.Get(s.domain)
			if err != nil || resp.StatusCode != 200 {
				s.status = 0
				s.statusLog = "DOWN"
			} else {
				s.status = 1
				s.statusLog = "UP"
			}
			log.Printf("%s %s\n", s.name, s.statusLog)
		}
		mutex.Unlock()
		time.Sleep(30 * time.Second)
	}
}
