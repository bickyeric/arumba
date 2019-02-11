package arumba

import (
	"log"
	"os"
	"strconv"

	"github.com/bickyeric/arumba/model"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	Instance Bot
)

type Bot struct {
	*tgbotapi.BotAPI
}

func ConfigureBot() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("telegramToken"))
	if err != nil {
		log.Fatal(err)
	}

	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	bot.Debug = debug

	Instance = Bot{bot}
}

func (bot Bot) SendHelpMessage(chatID int64) {
	tqMsg := tgbotapi.NewMessage(chatID, "Hai, coba deh klik /help")
	bot.Send(tqMsg)
}

func (bot Bot) SendNotFoundComic(chatID int64, comicName string) {
	tqMsg := tgbotapi.NewMessage(chatID, "Gk nemu nih bro comic +"+comicName+" ma :(")
	bot.Send(tqMsg)
}

func (bot Bot) SendNotFoundEpisode(chatID int64) {
	tqMsg := tgbotapi.NewMessage(chatID, "Gk nemu nih bro episode nya")
	bot.Send(tqMsg)
}

func (bot Bot) SendErrorMessage(chatID int64) {
	tqMsg := tgbotapi.NewMessage(chatID, "Waduh error nih bro maaf ya")
	bot.Send(tqMsg)
}

func (bot Bot) SendPage(chatID int64, pages []*model.Page) {
	tqMsg := tgbotapi.NewMessage(chatID, "Waduh error nih bro maaf ya")
	bot.Send(tqMsg)
}
