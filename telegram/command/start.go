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

type StartHandler struct {
	Bot    arumba.IBot
	Reader comic.Read
}

// Handle ...
func (s StartHandler) Handle(message *tgbotapi.Message) {
	arg := message.CommandArguments()

	if arg == "" {
		s.Bot.SendHelpMessage(message.Chat.ID)
		log.WithFields(
			log.Fields{
				"chat_id": message.Chat.ID,
			},
		).Info("Help message sent")
		return
	}

	comicName, episodeNo, err := s.parseArg(arg)
	if err != nil {
		s.Bot.SendHelpMessage(message.Chat.ID)
		return
	}

	pageURL, err := s.Reader.PerformByComicName(comicName, episodeNo)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			s.Bot.SendNotFoundEpisode(message.Chat.ID)
		default:
			s.Bot.SendErrorMessage(message.Chat.ID)
		}
		return
	}

	s.Bot.SendTextMessage(message.Chat.ID, pageURL)
}

func (s StartHandler) parseArg(arg string) (string, float64, error) {
	splittedString := strings.Split(arg, "_")
	episode, _ := strconv.ParseFloat(splittedString[1], 64)

	return splittedString[0], episode, nil
}
