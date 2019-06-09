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

type start struct {
	bot    arumba.IBot
	reader comic.Read
}

// NewStart ...
func NewStart(bot arumba.Bot, reader comic.Read) Handler {
	return start{bot, reader}
}

func (s start) Handle(message *tgbotapi.Message) {
	arg := message.CommandArguments()

	if arg == "" {
		s.bot.SendHelpMessage(message.Chat.ID)
		log.WithFields(
			log.Fields{
				"chat_id": message.Chat.ID,
			},
		).Info("Help message sent")
		return
	}

	comicName, episodeNo, err := s.parseArg(arg)
	if err != nil {
		s.bot.SendHelpMessage(message.Chat.ID)
		return
	}

	pageURL, err := s.reader.PerformByComicName(comicName, episodeNo)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			s.bot.SendNotFoundEpisode(message.Chat.ID)
		default:
			s.bot.SendErrorMessage(message.Chat.ID)
		}
		return
	}

	s.bot.SendTextMessage(message.Chat.ID, pageURL)
}

func (s start) parseArg(arg string) (string, float64, error) {
	splittedString := strings.Split(arg, "_")
	episode, _ := strconv.ParseFloat(splittedString[1], 64)

	return splittedString[0], episode, nil
}
