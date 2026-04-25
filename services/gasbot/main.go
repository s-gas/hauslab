package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	token, err := os.ReadFile("/run/secrets/telegram_token")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open %s\n", token)
		os.Exit(1)
	}
	for {
		fmt.Printf("%s\n", token)
		time.Sleep(1 * time.Second)
	}
}
