package callback

import (
	"errors"
	"strings"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/service/comic"
	"github.com/bickyeric/arumba/service/episode"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type handler struct {
	bot     arumba.IBot
	methods map[string]CallbackHandler
}

func NewHandler(app arumba.Arumba, bot arumba.IBot, kendang connection.IKendang) handler {
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
		Reader: comic.Read{
			ComicRepo:   app.ComicRepo,
			EpisodeRepo: app.EpisodeRepo,
			PageRepo:    app.PageRepo,
			Kendang:     kendang,
		},
	}
	return handler
}

func (handler handler) Handle(event *tgbotapi.CallbackQuery) {
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
