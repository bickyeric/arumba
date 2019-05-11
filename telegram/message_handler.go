package telegram

import (
	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/service/comic"
	"github.com/bickyeric/arumba/telegram/command"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

type messageHandler struct {
	startHandler    command.CommandHandler
	helpHandler     command.CommandHandler
	readHandler     command.CommandHandler
	feedbackHandler command.CommandHandler
	commonHandler   command.CommandHandler
}

func NewMessageHandler(app arumba.Arumba, bot arumba.IBot, kendang connection.IKendang) messageHandler {
	readerService := comic.Read{
		ComicRepo:   app.ComicRepo,
		EpisodeRepo: app.EpisodeRepo,
		PageRepo:    app.PageRepo,
		Kendang:     kendang,
	}
	return messageHandler{
		startHandler: command.StartHandler{
			Bot:    bot,
			Reader: readerService,
		},
		helpHandler: command.HelpHandler{
			Bot: bot,
		},
		readHandler: command.ReadHandler{
			Bot:    bot,
			Reader: readerService,
		},
		feedbackHandler: command.FeedbackHandler{
			Bot: bot,
		},
		commonHandler: command.CommonHandler{
			Bot: bot,
			ComicSearcher: comic.Search{
				Repo: app.ComicRepo,
			},
		},
	}
}

func (handler messageHandler) Handle(message *tgbotapi.Message) {
	log.WithFields(
		log.Fields{
			"text":    message.Text,
			"chat_id": message.Chat.ID,
		},
	).Info("Handling message")
	switch message.Command() {
	case command.StartCommand:
		go handler.startHandler.Handle(message)
	case command.HelpCommand:
		go handler.helpHandler.Handle(message)
	case command.ReadCommand:
		go handler.readHandler.Handle(message)
	case command.FeedbackCommand:
		go handler.feedbackHandler.Handle(message)
	default:
		go handler.commonHandler.Handle(message)
	}
}
