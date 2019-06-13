package telegram

import (
	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/service/comic"
	"github.com/bickyeric/arumba/service/episode"
	"github.com/bickyeric/arumba/telegram/message"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

type messageHandler map[string]message.Handler

// NewMessageHandler ...
func NewMessageHandler(app arumba.Arumba, bot arumba.Bot) message.Handler {
	readerService := comic.NewRead(app)
	finderService := comic.NewFinder(app.ComicRepo)

	handlers := messageHandler{}
	handlers[message.StartCommand] = message.NewStart(bot, readerService)
	handlers[message.HelpCommand] = message.NewHelp(bot)
	handlers[message.ReadCommand] = message.NewRead(bot, readerService, finderService, episode.NewSearch(app.EpisodeRepo))
	handlers[message.FeedbackCommand] = message.NewFeedback(bot)
	handlers[message.FollowCommand] = message.NewFollow(bot)
	handlers[message.GenericCommand] = message.NewGeneric(
		bot, bot,
		comic.NewSearch(app.ComicRepo),
	)

	return handlers
}

func (handler messageHandler) Handle(m *tgbotapi.Message) {
	log.WithFields(
		log.Fields{
			"text":    m.Text,
			"chat_id": m.Chat.ID,
		},
	).Info("Handling message")

	h, ok := handler[m.Command()]
	if ok {
		h.Handle(m)
	} else {
		handler[message.GenericCommand].Handle(m)
	}
}
