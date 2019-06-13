package main

import (
	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/telegram"
	log "github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
)

func main() {

	gotenv.Load(".env")
	log.SetFormatter(&log.JSONFormatter{})

	db := connection.NewMongo()

	bot := arumba.NewBot()

	app := arumba.New(db)

	log.Info("Webhook is running...")

	message := telegram.NewMessageHandler(app, bot)
	callback := telegram.NewCallbackHandler(app, bot)

	for update := range bot.UpdatesChannel() {
		if update.Message != nil {
			message.Handle(update.Message)

		} else if update.EditedMessage != nil {
			log.Info("received edited message event")

		} else if update.ChannelPost != nil {
			log.Info("received channel post event")

		} else if update.EditedChannelPost != nil {
			log.Info("received edited channel post event")

		} else if update.InlineQuery != nil {
			log.Info("received inline query event")

		} else if update.ChosenInlineResult != nil {
			log.Info("received chosen inline result event")

		} else if update.CallbackQuery != nil {
			callback.Handle(update.CallbackQuery)

		} else if update.ShippingQuery != nil {
			log.Info("received shipping query event")

		} else if update.PreCheckoutQuery != nil {
			log.Info("received pre checkout query event")
		}
	}
}
