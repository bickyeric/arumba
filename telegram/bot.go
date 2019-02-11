package telegram

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/bickyeric/arumba/model"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	BotInstance Bot
)

type Bot struct {
	*tgbotapi.BotAPI
}

func ConfigureBot() tgbotapi.UpdatesChannel {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("telegramToken"))
	if err != nil {
		log.Fatal(err)
	}

	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	bot.Debug = debug

	BotInstance = Bot{bot}

	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot.ListenForWebhook("/" + bot.Token)
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

type photoParams struct {
	ChatID int64  `json:"chat_id"`
	Photo  string `json:"photo"`
}

func (bot Bot) SendPage(chatID int64, pages []*model.Page) {
	url := "https://api.telegram.org/bot" + os.Getenv("telegramToken") + "/sendPhoto"
	for _, page := range pages {
		params := photoParams{chatID, page.Link}
		jsonParams, _ := json.Marshal(params)

		resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonParams))
		log.Print(resp.Body)
	}
}
