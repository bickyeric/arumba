package arumba

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bickyeric/arumba/model"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IBot ...
type IBot interface {
	SendReplyMessage(chatID int64, text string)
	SendTextMessage(chatID int64, text string) error
	SendComicSelector(chatID int64, comics []model.Comic)
	SendEpisodeSelector(chatID int64, comicID primitive.ObjectID, episodeGroup [][]float64)
	SendHelpMessage(chatID int64)
	SendNotFoundComic(chatID int64, comicName string)
	SendNotFoundEpisode(chatID int64)
	SendErrorMessage(chatID int64)

	NotifyError(err error)
	NotifyNewEpisode(model.Page)

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

func (bot bot) SendEpisodeSelector(chatID int64, comicID primitive.ObjectID, episodeGroup [][]float64) {
	tqMsg := tgbotapi.NewMessage(chatID, "OK. Select episode number below :D")
	keyboardRow := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{},
	}

	for _, group := range episodeGroup {
		data := ""
		text := ""
		if len(group) == 1 {
			data = fmt.Sprintf("select-episode_%s_%.1f", comicID.Hex(), group[0])
			text = fmt.Sprintf("%.1f", group[0])
		} else {
			data = fmt.Sprintf("select-episode_%s_%f_%f", comicID.Hex(), group[0], group[1])
			text = fmt.Sprintf("%.1f - %.1f", group[0], group[1])
		}

		keyboardRow.InlineKeyboard = append(keyboardRow.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(text, data),
		))
	}

	tqMsg.ReplyMarkup = keyboardRow
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

func (bot bot) SendTextMessage(chatID int64, text string) error {
	tqMsg := tgbotapi.NewMessage(chatID, text)
	_, err := bot.Send(tqMsg)
	return err
}

func (bot bot) SendComicSelector(chatID int64, comics []model.Comic) {
	tqMsg := tgbotapi.NewMessage(chatID, "Here we go, select comic below.")
	keyboardRow := [][]tgbotapi.InlineKeyboardButton{}

	for _, comic := range comics {
		data := fmt.Sprintf("select-comic_%s", comic.ID.Hex())
		keyboardRow = append(keyboardRow, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(comic.Name, data),
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
	chatID, _ := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 0)
	bot.SendTextMessage(chatID, "Error nih : "+err.Error())
}

func (bot bot) NotifyNewEpisode(page model.Page) {
	var tqMsg tgbotapi.MessageConfig
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	if debug {
		chatID, _ := strconv.Atoi(os.Getenv("CHAT_ID"))
		tqMsg = tgbotapi.NewMessage(int64(chatID), page.TelegraphLink)
	} else {
		tqMsg = tgbotapi.NewMessageToChannel(os.Getenv("TELEGRAM_CHANNEL"), page.TelegraphLink)
	}

	tqMsg.ParseMode = "Markdown"

	bot.Send(tqMsg)
}
