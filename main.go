package main

import (
	"flag"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

var token string

func init() {
	flag.StringVar(&token, "t", "", "telegram bot api token")
	flag.Parse()

	log.Println("token value:", token)

	if token == "" {
		log.Panic("No token was specified.")
	}
}

func main() {

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates, err := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "넵")
		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
}
