package main

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
	"os"
)

func main() {
	token := os.Getenv("TG_BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	updateCh, err := bot.GetUpdatesChan(ucfg)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case update := <-updateCh:
			userName := update.Message.From.UserName
			ChatID := update.Message.Chat.ID
			Text := update.Message.Text

			log.Printf("[%s] %d %s", userName, ChatID, Text)

			reply := Text
			msg := tgbotapi.NewMessage(ChatID, reply)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}
