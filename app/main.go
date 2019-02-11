package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/handler/telegram"
	"github.com/bickyeric/arumba/repository"
	"github.com/bickyeric/arumba/service"
	"github.com/subosito/gotenv"
)

var bot arumba.Bot

func main() {

	gotenv.Load(".env")

	connection.Connect()
	arumba.ConfigureBot()

	bot = arumba.Instance
	log.Printf("Webhook run on %s", os.Getenv("PORT"))
	log.Printf("Authorized on account %s", bot.Self.UserName)

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil)

	startHandler := telegram.Start{
		Bot: bot,
		ComicService: service.ComicService{
			ComicRepo:   repository.ComicRepository{},
			EpisodeRepo: repository.EpisodeRepository{},
			PageRepo:    repository.PageRepository{},
		},
	}

	for update := range updates {
		command := update.Message.Command()
		if command == "start" {
			startHandler.Handle(update.Message)
			// } else if command == "help" {
			// 	handler.HelpCommand(update.Message)
			// } else if command == "feedback" {
			// 	handler.FeedbackCommand(update.Message)
			// } else {
			// 	handler.Common(update.Message)
		}
	}
}
