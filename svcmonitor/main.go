package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	services := getServices()
	for {
		for _, s := range services {
			resp, err := http.Get(s.domain)
			if err != nil {
				fmt.Printf("%s is down\n", s.name)
			}
			if resp.StatusCode != 200 {
				fmt.Printf("%s is down\n", s.name)
			} else {
				fmt.Printf("%s is up\n", s.name)
			}
			time.Sleep(1 * time.Second)
		}
		time.Sleep(5 * time.Second)
	}
}
