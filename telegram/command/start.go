package command

import (
	"database/sql"
	"encoding/base64"
	"log"
	"strconv"
	"strings"

	"github.com/bickyeric/arumba/service"
	"github.com/bickyeric/arumba/telegram"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Start struct {
	Bot          telegram.Bot
	ComicService service.IComic
}

func (s Start) Handle(message *tgbotapi.Message) {
	arg := message.CommandArguments()

	if arg == "" {
		s.Bot.SendHelpMessage(message.Chat.ID)
		return
	}

	comicName, episodeNo := parseArg(arg)
	pages, err := s.ComicService.ReadComic(comicName, episodeNo)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			s.Bot.SendNotFoundEpisode(message.Chat.ID)
		default:
			s.Bot.SendErrorMessage(message.Chat.ID)
		}
		return
	}

	s.Bot.SendPage(message.Chat.ID, pages)
}

func parseArg(arg string) (string, float64) {
	decodedArg, err := base64.StdEncoding.DecodeString(arg)
	if err != nil {
		log.Fatal(err)
	}

	decodedString := string(decodedArg)
	splittedString := strings.Split(decodedString, "_")
	episode, _ := strconv.ParseFloat(splittedString[1], 64)

	return splittedString[0], episode
}