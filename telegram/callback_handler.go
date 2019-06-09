package telegram

import (
	"strings"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/service/comic"
	"github.com/bickyeric/arumba/service/episode"
	"github.com/bickyeric/arumba/service/telegraph"
	"github.com/bickyeric/arumba/telegram/callback"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

type CallbackHandler map[string]callback.Handler

func NewCallbackHandler(app arumba.Arumba, bot arumba.Bot, kendang connection.IKendang) CallbackHandler {
	telegraphCreator := telegraph.NewCreatePage()
	handlers := map[string]callback.Handler{}
	handlers[callback.SelectComicCallback] = callback.SelectComicHandler{
		Bot:      bot,
		Notifier: bot,
		EpisodeSearcher: episode.Search{
			Repo: app.EpisodeRepo,
		},
	}
	handlers[callback.SelectEpisodeCallback] = callback.SelectEpisodeHandler{
		Bot: bot,
		EpisodeSearcher: episode.Search{
			Repo: app.EpisodeRepo,
		},
		Reader: comic.NewRead(app, kendang, telegraphCreator),
	}
	return handlers
}

func (handler CallbackHandler) Handle(event *tgbotapi.CallbackQuery) {
	contextLog := log.WithFields(
		log.Fields{
			"data": event.Data,
		},
	)
	contextLog.Info("Handling callback")
	method, arg := handler.extractData(event.Data)
	h, ok := handler[method]

	if ok {
		h.Handle(event.Message.Chat.ID, arg)
	} else {
		contextLog.Warn("command not found : " + method)
	}
}

func (handler CallbackHandler) extractData(data string) (string, string) {
	arr := strings.Split(data, "_")

	return arr[0], data[len(arr[0])+1:]
}
