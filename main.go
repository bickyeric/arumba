package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("telegramToken"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe("0.0.0.0:"+port, nil)

	for update := range updates {
		log.Print("%+v\n", update)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Masih Development mang!!!")
		bot.Send(msg)
	}
}
