package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	b64 "encoding/base64"

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
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
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

func (bot Bot) NotifyError(err error) {
	tqMsg := tgbotapi.NewMessage(610339834, "Error nih : "+err.Error())
	bot.Send(tqMsg)
}

func (bot Bot) NotifyNewEpisode(update model.Update) {
	base64 := b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s_%f", update.ComicName, update.EpisodeNo)))
	txt := fmt.Sprintf("*%s*\nEpisode Baru!!!\nCek Sekarang juga :D!!!\n[klik disini](https://telegram.me/nb_comic_bot?start=%s)", update.ComicName, base64)
	tqMsg := tgbotapi.NewMessageToChannel(os.Getenv("TELEGRAM_CHANNEL"), txt)
	tqMsg.ParseMode = "Markdown"

	bot.Send(tqMsg)
}

func (bot Bot) SendPage(chatID int64, pages []*model.Page) {
	type photoParams struct {
		ChatID int64  `json:"chat_id"`
		Photo  string `json:"photo"`
	}

	url := "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_TOKEN") + "/sendPhoto"
	for _, page := range pages {
		params := photoParams{chatID, page.Link}
		jsonParams, _ := json.Marshal(params)

		resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonParams))
		log.Print(resp.Body)
	}
}
