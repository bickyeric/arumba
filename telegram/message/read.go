package message

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/service/comic"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

var comicNameRequest = "OK. You want to read a comic, just give me a comic name."

type read struct {
	bot    arumba.IBot
	reader comic.Read
}

// NewRead ...
func NewRead(bot arumba.Bot, reader comic.Read) Handler {
	return read{bot, reader}
}

func (r read) Handle(message *tgbotapi.Message) {
	arg := message.CommandArguments()
	comicName, episodeNo := r.parseArg(arg)

	if comicName != "" && episodeNo > -1 {
		r.readComicEpisode(message.Chat.ID, comicName, episodeNo)
	} else if comicName != "" {
		r.readComic(message.Chat.ID, comicName)
	} else {
		r.requestComicName(message.Chat.ID)
	}
}

func (r read) requestComicName(chatID int64) {
	r.bot.SendReplyMessage(chatID, comicNameRequest)
	log.WithFields(
		log.Fields{
			"chat_id": chatID,
		},
	).Info("Request comic name message sent")
}

func (r read) readComic(chatID int64, comicName string) {
	log.Println("ada nama_comic saja")
}

func (r read) readComicEpisode(chatID int64, comicName string, episodeNo float64) {
	pageURL, err := r.reader.PerformByComicName(comicName, episodeNo)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			r.bot.SendNotFoundEpisode(chatID)
		default:
			r.bot.SendErrorMessage(chatID)
		}
		return
	}

	r.bot.SendTextMessage(chatID, pageURL)
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
