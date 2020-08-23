package main

import (
	"github.com/DmitriyShavkunov/streamdj-tgbot/internal/streamdj"
	"log"
	"os"
)

func main() {
	baseURL := os.Getenv("STREAMDJ_API_BASE_URL")
	channelID := os.Getenv("STREAMDJ_CHANNEL_ID")
	apiKey := os.Getenv("STREAMDJ_API_KEY")
	err := streamdj.SkipTrack(baseURL, channelID, apiKey)
	if err != nil {
		panic(err)
	}
	log.Print("success request")
}
