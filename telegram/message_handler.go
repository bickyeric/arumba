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

type MessageHandler map[string]message.Handler

// NewMessageHandler ...
func NewMessageHandler(app arumba.Arumba, bot arumba.IBot, kendang connection.IKendang) MessageHandler {
	telegraphCreator := telegraph.NewCreatePage()
	readerService := comic.NewRead(app, kendang, telegraphCreator)
	handlers := map[string]message.Handler{}

	handlers[message.StartCommand] = message.StartHandler{
		Bot:    bot,
		Reader: readerService,
	}

	handlers[message.HelpCommand] = message.HelpHandler{
		Bot: bot,
	}

	handlers[message.ReadCommand] = message.ReadHandler{
		Bot:    bot,
		Reader: readerService,
	}

	handlers[message.FeedbackCommand] = message.FeedbackHandler{
		Bot: bot,
	}

	handlers[message.FollowCommand] = message.FollowHandler{
		Bot: bot,
	}

	handlers[message.GenericCommand] = message.GenericHandler{
		Bot: bot,
		ComicSearcher: comic.Search{
			Repo: app.ComicRepo,
		},
	}

	return handlers
}

func (handler MessageHandler) Handle(m *tgbotapi.Message) {
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
