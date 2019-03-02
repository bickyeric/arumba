package command

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	"github.com/bickyeric/arumba/service/comic"
	"github.com/bickyeric/arumba/telegram"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var comicNameRequest = "OK. You want to read a comic, just give me a comic name."

type read struct {
	bot    telegram.IBot
	reader comic.Read
}

func Read(bot telegram.IBot, reader comic.Read) telegram.CommandHandler {
	return read{bot, reader}
}

func (r read) Handle(message *tgbotapi.Message) {
	arg := message.CommandArguments()
	comicName, episodeNo := r.parseArg(arg)

	if comicName != "" && episodeNo > -1 {
		r.readComicEpisode(message.Chat.ID, comicName, episodeNo)
	} else if comicName != "" {
		log.Println("ada nama_comic saja")
	} else {
		r.requestComicName(message.Chat.ID)
	}
}

func (r read) requestComicName(chatID int64) {
	r.bot.SendReplyMessage(chatID, comicNameRequest)
}

func (r read) readComicEpisode(chatID int64, comicName string, episodeNo float64) {
	pages, err := r.reader.Perform(comicName, episodeNo)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			r.bot.SendNotFoundEpisode(chatID)
		default:
			r.bot.SendErrorMessage(chatID)
		}
		return
	}

	r.bot.SendPage(chatID, pages)
}

func (r read) parseArg(arg string) (string, float64) {
	if arg == "" {
		return "", -1
	}

	words := strings.Split(arg, " ")
	episodeNo, err := strconv.ParseFloat(words[len(words)-1], 64)
	if err != nil {
		return arg, -1
	}
	comicName := strings.Join(words[:len(words)-1], " ")

	return comicName, episodeNo
}
