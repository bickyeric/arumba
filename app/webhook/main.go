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
	bot := arumba.NewBot()
	kendang := connection.NewKendang()

	app := arumba.New(db)

	log.Printf("Webhook run on %s", os.Getenv("PORT"))
	go http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil)

	messageHandler := telegram.NewMessageHandler(app, bot, kendang)
	callbackHandler := telegram.NewCallbackHandler(app, bot)

	for update := range bot.UpdatesChannel() {
		if update.Message != nil {
			messageHandler.Handle(update.Message)
			continue
		}
		if update.EditedMessage != nil {
			log.Println("received edited message event")
			continue
		}
		if update.ChannelPost != nil {
			log.Println("received channel post event")
			continue
		}
		if update.EditedChannelPost != nil {
			log.Println("received edited channel post event")
			continue
		}
		if update.InlineQuery != nil {
			log.Println("received inline query event")
			continue
		}
		if update.ChosenInlineResult != nil {
			log.Println("received chosen inline result event")
			continue
		}
		if update.CallbackQuery != nil {
			callbackHandler.Handle(update.CallbackQuery)
			continue
		}
		if update.ShippingQuery != nil {
			log.Println("received shipping query event")
			continue
		}
		if update.PreCheckoutQuery != nil {
			log.Println("received pre checkout query event")
			continue
		}
	}
}
