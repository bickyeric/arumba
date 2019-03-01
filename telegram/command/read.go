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

// Read ...
type Read struct {
	Bot    telegram.Bot
	Reader comic.Read
}

// Handle ...
func (r Read) Handle(message *tgbotapi.Message) {
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

func (r Read) requestComicName(chatID int64) {
	r.Bot.SendReplyMessage(chatID, comicNameRequest)
}

func (r Read) readComicEpisode(chatID int64, comicName string, episodeNo float64) {
	pages, err := r.Reader.Perform(comicName, episodeNo)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			r.Bot.SendNotFoundEpisode(chatID)
		default:
			r.Bot.SendErrorMessage(chatID)
		}
		return
	}

	r.Bot.SendPage(chatID, pages)
}

func (r Read) parseArg(arg string) (string, float64) {
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
