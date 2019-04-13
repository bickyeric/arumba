package callback

import (
	"errors"
	"log"
	"strings"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/service/episode"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type handler struct {
	bot     arumba.IBot
	methods map[string]CallbackHandler
}

func NewHandler(app arumba.Arumba, bot arumba.IBot) handler {
	handler := handler{
		bot:     bot,
		methods: map[string]CallbackHandler{},
	}
	handler.methods[ReadCallback] = ReadHandler{
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
	}
	return handler
}

func (handler handler) Handle(event *tgbotapi.CallbackQuery) {
	method, arg := handler.extractData(event.Data)
	log.Println(method, arg)
	h, ok := handler.methods[method]

	if ok {
		h.Handle(event.Message.Chat.ID, arg)
	} else {
		handler.bot.NotifyError(errors.New("command not found : " + method))
	}
}

func (handler handler) extractData(data string) (string, string) {
	arr := strings.Split(string(data), "_")

	return arr[0], arr[1]
}
