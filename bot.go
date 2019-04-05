package arumba

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

// IBot ...
type IBot interface {
	SendReplyMessage(chatID int64, text string)
	SendTextMessage(chatID int64, text string)
	SendComicSelector(chatID int64, comics []model.Comic)
	SendEpisodeSelector(chatID int64, comicID int, episodeGroup [][]float64)
	SendHelpMessage(chatID int64)
	SendNotFoundComic(chatID int64, comicName string)
	SendNotFoundEpisode(chatID int64)
	SendErrorMessage(chatID int64)

	NotifyError(err error)
	NotifyNewEpisode(update model.Update)
	SendPage(chatID int64, links []string)

	Bot() bot
	UpdatesChannel() tgbotapi.UpdatesChannel
}

type bot struct {
	*tgbotapi.BotAPI
}

// NewBot ...
func NewBot() IBot {
	botapi, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	botapi.Debug = false

	return bot{botapi}
}

func (bot bot) SendEpisodeSelector(chatID int64, comicID int, episodeGroup [][]float64) {
	tqMsg := tgbotapi.NewMessage(chatID, "OK. Select episode number below :D")
	keyboardRow := [][]tgbotapi.InlineKeyboardButton{}

	for _, group := range episodeGroup {
		base64 := ""
		text := ""
		if len(group) == 1 {
			base64 = b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("select-episode_%d_%f", comicID, group[0])))
			text = fmt.Sprintf("%.1f", group[0])
		} else {
			base64 = b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("select-episode_%d_%f_%f", comicID, group[0], group[1])))
			text = fmt.Sprintf("%.1f - %.1f", group[0], group[1])
		}

		keyboardRow = append(keyboardRow, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(text, base64),
		))
	}

	tqMsg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(keyboardRow...)
	bot.Send(tqMsg)
}

func (bot bot) UpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)
	return updates
}

func (bot bot) Bot() bot {
	return bot
}

func (bot bot) SendReplyMessage(chatID int64, text string) {
	replyMsg := tgbotapi.NewMessage(chatID, text)
	replyMsg.ReplyMarkup = tgbotapi.ForceReply{ForceReply: true}
	bot.Send(replyMsg)
}

func (bot bot) SendTextMessage(chatID int64, text string) {
	tqMsg := tgbotapi.NewMessage(chatID, text)
	bot.Send(tqMsg)
}

func (bot bot) SendComicSelector(chatID int64, comics []model.Comic) {
	tqMsg := tgbotapi.NewMessage(chatID, "Here we go, select comic below.")
	keyboardRow := [][]tgbotapi.InlineKeyboardButton{}

	for _, comic := range comics {
		base64 := b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("read_%d", comic.ID)))
		keyboardRow = append(keyboardRow, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(comic.Name, base64),
		))
	}

	tqMsg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(keyboardRow...)
	bot.Send(tqMsg)
}

func (bot bot) SendHelpMessage(chatID int64) {
	bot.SendTextMessage(chatID, "Hai, coba deh klik /help")
}

func (bot bot) SendNotFoundComic(chatID int64, comicName string) {
	bot.SendTextMessage(chatID, "Gk nemu nih bro comic "+comicName+" ma :(")
}

func (bot bot) SendNotFoundEpisode(chatID int64) {
	bot.SendTextMessage(chatID, "Gk nemu nih bro episode nya")
}

func (bot bot) SendErrorMessage(chatID int64) {
	bot.SendTextMessage(chatID, "Waduh error nih bro maaf ya")
}

func (bot bot) NotifyError(err error) {
	chatID, _ := strconv.ParseInt(os.Getenv("CHAT_ID"), 36, 0)
	bot.SendTextMessage(chatID, "Error nih : "+err.Error())
}

func (bot bot) NotifyNewEpisode(update model.Update) {
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

func (bot bot) SendPage(chatID int64, links []string) {
	type photoParams struct {
		ChatID int64  `json:"chat_id"`
		Photo  string `json:"photo"`
	}

	url := "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_TOKEN") + "/sendPhoto"
	for _, link := range links {
		params := photoParams{chatID, link}
		jsonParams, _ := json.Marshal(params)

		http.Post(url, "application/json", bytes.NewBuffer(jsonParams))
	}
}
