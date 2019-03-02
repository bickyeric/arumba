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

type start struct {
	bot    telegram.IBot
	reader comic.Read
}

func Start(bot telegram.IBot, reader comic.Read) telegram.CommandHandler {
	return start{bot, reader}
}

// Handle ...
func (s start) Handle(message *tgbotapi.Message) {
	arg := message.CommandArguments()

	if arg == "" {
		s.bot.SendHelpMessage(message.Chat.ID)
		return
	}

	comicName, episodeNo, err := s.parseArg(arg)
	if err != nil {
		s.bot.SendHelpMessage(message.Chat.ID)
		return
	}

	pages, err := s.reader.Perform(comicName, episodeNo)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			s.bot.SendNotFoundEpisode(message.Chat.ID)
		default:
			s.bot.SendErrorMessage(message.Chat.ID)
		}
		return
	}

	s.bot.SendPage(message.Chat.ID, pages)
}

func (s start) parseArg(arg string) (string, float64, error) {
	decodedArg, err := base64.StdEncoding.DecodeString(arg)
	if err != nil {
		return "", 0.0, err
	}

	decodedString := string(decodedArg)
	splittedString := strings.Split(decodedString, "_")
	episode, _ := strconv.ParseFloat(splittedString[1], 64)

	return splittedString[0], episode, nil
}
