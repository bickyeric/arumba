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

	connection.Connect()
	arumba.Configure()
	updates := telegram.ConfigureBot()

	app := arumba.Instance
	startHandler := app.InjectTelegramStart()
	helpHandler := app.InjectTelegramHelp()
	readHandler := app.InjectTelegramRead()

	log.Printf("Webhook run on %s", os.Getenv("PORT"))
	go http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil)

	for update := range updates {
		command := update.Message.Command()
		if command == "start" {
			startHandler.Handle(update.Message)
		} else if command == "help" {
			helpHandler.Handle(update.Message)
		} else if command == "read" {
			readHandler.Handle(update.Message)
			// } else if command == "feedback" {
			// 	handler.FeedbackCommand(update.Message)
			// } else {
			// 	handler.Common(update.Message)
		}
	}
}
