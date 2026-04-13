package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello from svcmonitor")
	resp, err := http.Get("http://localhost:3000")
	if err != nil {
		fmt.Println("Dead")
		return
	}
	fmt.Println("Alive", resp)
}
