package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	for {
		resp, err := http.Get("http://localhost:3000")
		if err != nil {
			fmt.Println("Server listening at 3000 is down")
		}
		if resp.StatusCode != 200 {
			fmt.Println("Server listening at 3000 is down")
		} else {
			fmt.Println("Server listening at 3000 is up")
		}
		time.Sleep(5 * time.Second)
	}
}
