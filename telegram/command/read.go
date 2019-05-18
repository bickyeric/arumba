package command

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

type ReadHandler struct {
	Bot    arumba.IBot
	Reader comic.Read
}

func (r ReadHandler) Handle(message *tgbotapi.Message) {
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

func (r ReadHandler) requestComicName(chatID int64) {
	r.Bot.SendReplyMessage(chatID, comicNameRequest)
	log.WithFields(
		log.Fields{
			"chat_id": chatID,
		},
	).Info("Request comic name message sent")
}

func (r ReadHandler) readComicEpisode(chatID int64, comicName string, episodeNo float64) {
	pageURL, err := r.Reader.PerformByComicName(comicName, episodeNo)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			r.Bot.SendNotFoundEpisode(chatID)
		default:
			r.Bot.SendErrorMessage(chatID)
		}
		return
	}

	r.Bot.SendTextMessage(chatID, pageURL)
}

func (r ReadHandler) parseArg(arg string) (string, float64) {
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
