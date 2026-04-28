package main

import (
	"log"
	"net/http"
	"os"
)

type service struct {
	name   string
	domain string
	port   int
	status int
}

func getServices() map[string]service {
	services := make(map[string]service)
	services["sysmetrics"] = service{
		name:   "sysmetrics",
		port:   1024,
		domain: "http://sysmetrics:1024/metrics",
		status: 0,
	}
	services["prometheus"] = service{
		name:   "prometheus",
		port:   9090,
		domain: "http://prometheus:9090",
		status: 0,
	}
	services["grafana"] = service{
		name:   "grafana",
		port:   3000,
		domain: "http://grafana:3000",
		status: 0,
	}
	return services
}

func checkServices(services map[string]service, logFile *os.File) {
	log.SetOutput(logFile)
	for _, s := range services {
		resp, err := http.Get(s.domain)
		if err != nil || resp.StatusCode != 200 {
			log.Printf("%s DOWN\n", s.name)
			s.status = 0
		} else {
			log.Printf("%s UP\n", s.name)
			s.status = 1
		}
	}
}
