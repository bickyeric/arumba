package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

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

	bot.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	BotInstance = Bot{bot}
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	log.Printf("Authorized on account %s", bot.Self.UserName)
	return updates
}

func (bot Bot) SendReplyMessage(chatID int64, text string) {
	replyMsg := tgbotapi.NewMessage(chatID, text)
	replyMsg.ReplyMarkup = tgbotapi.ForceReply{ForceReply: true}
	bot.Send(replyMsg)
}

func (bot Bot) SendTextMessage(chatID int64, text string) {
	tqMsg := tgbotapi.NewMessage(chatID, text)
	bot.Send(tqMsg)
}

func (bot Bot) SendComicSelector(chatID int64, comics []model.Comic) {
	tqMsg := tgbotapi.NewMessage(chatID, "Here we go, select comic below.")
	keyboardRow := [][]tgbotapi.InlineKeyboardButton{}

	for _, comic := range comics {
		keyboardRow = append(keyboardRow, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(comic.Name, fmt.Sprintf("comicID_%d", comic.ID)),
		))
	}

	tqMsg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(keyboardRow...)
	_, err := bot.Send(tqMsg)
	log.Println(err)
}

func (bot Bot) SendHelpMessage(chatID int64) {
	bot.SendTextMessage(chatID, "Hai, coba deh klik /help")
}

func (bot Bot) SendNotFoundComic(chatID int64, comicName string) {
	bot.SendTextMessage(chatID, "Gk nemu nih bro comic +"+comicName+" ma :(")
}

func (bot Bot) SendNotFoundEpisode(chatID int64) {
	bot.SendTextMessage(chatID, "Gk nemu nih bro episode nya")
}

func (bot Bot) SendErrorMessage(chatID int64) {
	bot.SendTextMessage(chatID, "Waduh error nih bro maaf ya")
}

func (bot Bot) NotifyError(err error) {
	chatID, _ := strconv.ParseInt(os.Getenv("CHAT_ID"), 36, 0)
	bot.SendTextMessage(chatID, "Error nih : "+err.Error())
}

func (bot Bot) NotifyNewEpisode(update model.Update) {
	base64 := b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s_%f", update.ComicName, update.EpisodeNo)))
	txt := fmt.Sprintf("*%s*\nEpisode Baru!!!\nCek Sekarang juga :D!!!\n[klik disini](https://telegram.me/nb_comic_bot?start=%s)", update.ComicName, base64)

	var tqMsg tgbotapi.MessageConfig
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	if debug {
		chatID, _ := strconv.Atoi(os.Getenv("CHAT_ID"))
		tqMsg = tgbotapi.NewMessage(int64(chatID), txt)
	} else {
		tqMsg = tgbotapi.NewMessageToChannel(os.Getenv("TELEGRAM_CHANNEL"), txt)
	}

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

		http.Post(url, "application/json", bytes.NewBuffer(jsonParams))
	}
}
