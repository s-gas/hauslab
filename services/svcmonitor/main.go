package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Services map[string]*Service `yaml:"services"`
}

func main() {
	data, err := os.ReadFile("services.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	yaml.Unmarshal(data, &config)
	services := config.Services

	endpoint := "/metrics"
	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		serveMetrics(w, services)
	})

	go checkServices(services)

	port := 1024
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("svcmonitor listening at %s%s\n", addr, endpoint)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
