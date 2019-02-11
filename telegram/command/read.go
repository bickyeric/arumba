package command

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	"github.com/bickyeric/arumba/service"
	"github.com/bickyeric/arumba/telegram"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Read struct {
	Bot          telegram.Bot
	ComicService service.IComic
}

func (r Read) Handle(message *tgbotapi.Message) {
	arg := message.CommandArguments()
	comicName, episodeNo := r.parseArg(arg)

	if comicName != "" && episodeNo > -1 {
		// log.Println("ada nama_comic dan episodenya")
		r.readComicEpisode(message.Chat.ID, comicName, episodeNo)
	} else if comicName != "" {
		log.Println("ada nama_comic saja")
	} else {
		log.Println("pengen baca aja")
	}
}

func (r Read) readComicEpisode(chatID int64, comicName string, episodeNo int) {
	pages, err := r.ComicService.ReadComic(comicName, episodeNo)
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

func (r Read) parseArg(arg string) (string, int) {
	if arg == "" {
		return "", -1
	}

	words := strings.Split(arg, " ")
	episodeNo, err := strconv.Atoi(words[len(words)-1])
	if err != nil {
		return arg, -1
	}
	comicName := strings.Join(words[:len(words)-1], " ")

	return comicName, episodeNo
}
