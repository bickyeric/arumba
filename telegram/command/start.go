package command

import (
	"database/sql"
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/bickyeric/arumba/service/comic"
	"github.com/bickyeric/arumba/telegram"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// Start ...
type Start struct {
	Bot    telegram.Bot
	Reader comic.Read
}

// Handle ...
func (s Start) Handle(message *tgbotapi.Message) {
	arg := message.CommandArguments()

	if arg == "" {
		s.Bot.SendHelpMessage(message.Chat.ID)
		return
	}

	comicName, episodeNo, err := s.parseArg(arg)
	if err != nil {
		s.Bot.SendHelpMessage(message.Chat.ID)
		return
	}

	pages, err := s.Reader.Perform(comicName, episodeNo)
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

func (s Start) parseArg(arg string) (string, float64, error) {
	decodedArg, err := base64.StdEncoding.DecodeString(arg)
	if err != nil {
		return "", 0.0, err
	}

	decodedString := string(decodedArg)
	splittedString := strings.Split(decodedString, "_")
	episode, _ := strconv.ParseFloat(splittedString[1], 64)

	return splittedString[0], episode, nil
}
