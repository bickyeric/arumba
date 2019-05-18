package callback

import (
	"errors"
	"strings"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/service/comic"
	"github.com/bickyeric/arumba/service/episode"
	"github.com/bickyeric/arumba/service/telegraph"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

type handler struct {
	bot     arumba.IBot
	methods map[string]CallbackHandler
}

func NewHandler(app arumba.Arumba, bot arumba.IBot, kendang connection.IKendang) handler {
	telegraphCreator := telegraph.NewCreatePage()
	handler := handler{
		bot:     bot,
		methods: map[string]CallbackHandler{},
	}
	handler.methods[SelectComicCallback] = SelectComicHandler{
		Bot: bot,
		EpisodeSearcher: episode.Search{
			Repo: app.EpisodeRepo,
		},
	}
	handler.methods[SelectEpisodeCallback] = SelectEpisodeHandler{
		Bot: bot,
		EpisodeSearcher: episode.Search{
			Repo: app.EpisodeRepo,
		},
		Reader: comic.NewRead(app, kendang, telegraphCreator),
	}
	return handler
}

func (handler handler) Handle(event *tgbotapi.CallbackQuery) {
	log.WithFields(
		log.Fields{
			"data": event.Data,
		},
	).Info("Handling callback")
	method, arg := handler.extractData(event.Data)
	h, ok := handler.methods[method]

	if ok {
		h.Handle(event.Message.Chat.ID, arg)
	} else {
		handler.bot.NotifyError(errors.New("command not found : " + method))
	}
}

func (handler handler) extractData(data string) (string, string) {
	arr := strings.Split(data, "_")

	return arr[0], data[len(arr[0])+1:]
}
