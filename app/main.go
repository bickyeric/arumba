package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/handler"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot tgbotapi.BotAPI

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	log.Printf("Webhook run on %s", port)

	bot := arumba.Instance()
	log.Printf("Authorized on account %s", bot.Self.UserName)

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe("0.0.0.0:"+port, nil)

	for update := range updates {
		command := update.Message.Command()
		if command == "start" {
			handler.StartCommand(update.Message)
		} else if command == "help" {
			handler.HelpCommand(update.Message)
		} else if command == "feedback" {
			handler.FeedbackCommand(update.Message)
		} else {
			handler.Common(update.Message)
		}
	}
}
