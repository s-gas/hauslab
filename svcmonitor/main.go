package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	services := getServices()
	s := services["sysmetrics"]
	for {
		resp, err := http.Get(s.domain)
		if err != nil {
			fmt.Printf("%s:%d is down\n", s.name, s.port)
		}
		if resp.StatusCode != 200 {
			fmt.Printf("%s:%d is down\n", s.name, s.port)
		} else {
			fmt.Printf("%s:%d is up\n", s.name, s.port)
		}
		time.Sleep(5 * time.Second)
	}
}
