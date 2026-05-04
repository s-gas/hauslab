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
	mutex     sync.Mutex
}

func getServices() map[string]*service {
	services := make(map[string]*service)
	services["sysmetrics"] = &service{
		name:   "sysmetrics",
		port:   1024,
		domain: "http://sysmetrics:1024/metrics",
		status: 0,
	}
	services["prometheus"] = &service{
		name:   "prometheus",
		port:   9090,
		domain: "http://prometheus:9090",
		status: 0,
	}
	services["grafana"] = &service{
		name:   "grafana",
		port:   3000,
		domain: "http://grafana:3000",
		status: 0,
	}
	services["adguardhome"] = &service{
		name:   "adguardhome",
		port:   80,
		domain: "http://adguardhome:80",
		status: 0,
	}
	services["nginx"] = &service{
		name:   "nginx",
		port:   80,
		domain: "http://nginx:80",
		status: 0,
	}
	return services
}

func checkServices(services map[string]*service) {
	for {
		for _, s := range services {
			checkService(s)
		}
		time.Sleep(30 * time.Second)
	}
}

func checkService(s *service) {
	resp, err := http.Get(s.domain)
	if err == nil {
		defer resp.Body.Close()
	}
	s.mutex.Lock()
	if err != nil || resp.StatusCode != 200 {
		s.status = 0
		log.Printf("%s DOWN\n", s.name)
	} else {
		s.status = 1
		log.Printf("%s UP\n", s.name)
	}
	s.mutex.Unlock()
}
