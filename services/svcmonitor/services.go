package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

type Service struct {
	Url       string `yaml:"domain"`
	statusLog string
	Port      int `yaml:"port"`
	status    int
	mutex     sync.Mutex
}

func checkServices(services map[string]*Service) {
	for {
		for name, s := range services {
			checkService(name, s)
		}
		time.Sleep(30 * time.Second)
	}
}

func checkService(name string, s *Service) {
	resp, err := http.Get(s.Url)
	if err == nil {
		defer resp.Body.Close()
	}
	s.mutex.Lock()
	if err != nil || resp.StatusCode != 200 {
		s.status = 0
		log.Printf("%s DOWN\n", name)
	} else {
		s.status = 1
		log.Printf("%s UP\n", name)
	}
	s.mutex.Unlock()
}
