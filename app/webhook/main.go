package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/telegram"
	"github.com/bickyeric/arumba/telegram/callback"
	"github.com/subosito/gotenv"
)

func main() {

	gotenv.Load(".env")

	db, err := connection.NewMongo()
	if err != nil {
		log.Fatal(err)
	}
	bot := arumba.NewBot()
	kendang := connection.NewKendang()

	app := arumba.New(db)

	log.Printf("Webhook run on %s", os.Getenv("PORT"))
	go http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil)

	messageHandler := telegram.NewMessageHandler(app, bot, kendang)
	callback := callback.NewHandler(app, bot, kendang)

	for update := range bot.UpdatesChannel() {
		if update.Message != nil {
			messageHandler.Handle(update.Message)

		} else if update.EditedMessage != nil {
			log.Println("received edited message event")

		} else if update.ChannelPost != nil {
			log.Println("received channel post event")

		} else if update.EditedChannelPost != nil {
			log.Println("received edited channel post event")

		} else if update.InlineQuery != nil {
			log.Println("received inline query event")

		} else if update.ChosenInlineResult != nil {
			log.Println("received chosen inline result event")

		} else if update.CallbackQuery != nil {
			callback.Handle(update.CallbackQuery)

		} else if update.ShippingQuery != nil {
			log.Println("received shipping query event")

		} else if update.PreCheckoutQuery != nil {
			log.Println("received pre checkout query event")
		}
	}
}
