package main

import (
	"context"
	"os"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/telegram"
	log "github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
)

func main() {
	ctx := context.Background()

	gotenv.Load(".env")
	log.SetFormatter(&log.JSONFormatter{})

	db := connection.NewMongo(ctx).Database(os.Getenv("DB_MONGO_DATABASE"))

	bot := arumba.NewBot()

	app := arumba.New(db)

	log.Info("Webhook is running...")

	message := telegram.NewMessageHandler(app, bot)
	callback := telegram.NewCallbackHandler(app, bot)

	for update := range bot.UpdatesChannel() {
		if update.Message != nil {
			message.Handle(update.Message)
		}
		if update.EditedMessage != nil {
			log.Info("received edited message event")
		}
		if update.ChannelPost != nil {
			log.Info("received channel post event")
		}
		if update.EditedChannelPost != nil {
			log.Info("received edited channel post event")
		}
		if update.InlineQuery != nil {
			log.Info("received inline query event")
		}
		if update.ChosenInlineResult != nil {
			log.Info("received chosen inline result event")
		}
		if update.CallbackQuery != nil {
			callback.Handle(update.CallbackQuery)
		}
		if update.ShippingQuery != nil {
			log.Info("received shipping query event")
		}
		if update.PreCheckoutQuery != nil {
			log.Info("received pre checkout query event")
		}
	}
}
