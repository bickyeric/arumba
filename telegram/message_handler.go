package telegram

import (
	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/service/comic"
	"github.com/bickyeric/arumba/telegram/command"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type messageHandler struct {
	startHandler    CommandHandler
	helpHandler     CommandHandler
	readHandler     CommandHandler
	feedbackHandler CommandHandler
	commonHandler   CommandHandler
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
	switch message.Command() {
	case StartCommand:
		go handler.startHandler.Handle(message)
	case HelpCommand:
		go handler.helpHandler.Handle(message)
	case ReadCommand:
		go handler.readHandler.Handle(message)
	case FeedbackCommand:
		go handler.feedbackHandler.Handle(message)
	default:
		go handler.commonHandler.Handle(message)
	}
}
