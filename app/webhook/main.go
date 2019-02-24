package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/telegram"
	"github.com/subosito/gotenv"
)

func main() {

	gotenv.Load(".env")

	db := connection.NewMysql()
	updates := telegram.ConfigureBot()

	app := arumba.New(telegram.BotInstance, db)
	startHandler := app.InjectTelegramStart()
	helpHandler := app.InjectTelegramHelp()
	readHandler := app.InjectTelegramRead()

	log.Printf("Webhook run on %s", os.Getenv("PORT"))
	go http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil)

	for update := range updates {
		switch update.Message.Command() {
		case telegram.StartCommand:
			go startHandler.Handle(update.Message)
		case telegram.HelpCommand:
			go helpHandler.Handle(update.Message)
		case telegram.ReadCommand:
			go readHandler.Handle(update.Message)
			// case telegram.FeedbackCommand:
			// 	handler.FeedbackCommand(update.Message)
			// default:
			// 	handler.Common(update.Message)
		}
	}
}
