package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	secret, err := os.ReadFile("/run/secrets/telegram_token")
	if err != nil {
		fmt.Printf("error\n")
		os.Exit(1)
	}
	for {
		fmt.Printf("%s\n", secret)
		time.Sleep(1 * time.Second)
	}
}
