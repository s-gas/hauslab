package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/s-gas/hauslab/services/gasmetrics/gasmetrics-bot/internal/reading"
	"github.com/s-gas/hauslab/services/gasmetrics/gasmetrics-bot/internal/telegram"
	"log"
	"net/http"
	"os"
)

var url = "http://gasmetrics-server" // Docker domain name

func main() {
	bot, err := telegram.CreateBot()
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	url := fmt.Sprintf("%s:%s", url, port)
	contentType := "application/json"

	updates := telegram.GetUpdates(bot)

	for update := range updates {
		var msg string
		r, err := reading.Parse(update.Message.Text)
		if err != nil {
			log.Println(err)
			msg = fmt.Sprintf("Failed to add entry: %v", err)
			telegram.Reply(msg, bot, update)
			continue
		}
		body, err := json.Marshal(r)
		if err != nil {
			log.Println(err)
			continue
		}
		resp, err := http.Post(url, contentType, bytes.NewReader(body))
		if err != nil {
			log.Println(err)
			continue
		}
		resp.Body.Close()

		if resp.StatusCode == http.StatusCreated {
			msg = "Entry added successfully"
		} else {
			msg = "Failed to add entry: invalid input"
		}
		log.Println(msg)
		telegram.Reply(msg, bot, update)
	}
}
