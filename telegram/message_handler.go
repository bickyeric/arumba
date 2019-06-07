package telegram

import (
	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/service/comic"
	"github.com/bickyeric/arumba/service/telegraph"
	"github.com/bickyeric/arumba/telegram/message"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

type messageHandler struct {
	commands map[string]message.Handler
}

func NewMessageHandler(app arumba.Arumba, bot arumba.IBot, kendang connection.IKendang) messageHandler {
	telegraphCreator := telegraph.NewCreatePage()
	readerService := comic.NewRead(app, kendang, telegraphCreator)
	handler := messageHandler{
		commands: map[string]message.Handler{},
	}

	handler.commands[message.StartCommand] = message.StartHandler{
		Bot:    bot,
		Reader: readerService,
	}

	handler.commands[message.HelpCommand] = message.HelpHandler{
		Bot: bot,
	}

	handler.commands[message.ReadCommand] = message.ReadHandler{
		Bot:    bot,
		Reader: readerService,
	}

	handler.commands[message.FeedbackCommand] = message.FeedbackHandler{
		Bot: bot,
	}

	handler.commands[message.GenericCommand] = message.GenericHandler{
		Bot: bot,
		ComicSearcher: comic.Search{
			Repo: app.ComicRepo,
		},
	}

	return handler
}

func (handler messageHandler) Handle(m *tgbotapi.Message) {
	log.WithFields(
		log.Fields{
			"text":    m.Text,
			"chat_id": m.Chat.ID,
		},
	).Info("Handling message")

	h, ok := handler.commands[m.Command()]
	if ok {
		h.Handle(m)
	} else {
		handler.commands[message.GenericCommand].Handle(m)
	}
}
