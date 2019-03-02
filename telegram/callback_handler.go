package telegram

import (
	"encoding/base64"
	"errors"
	"strings"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/service/episode"
	"github.com/bickyeric/arumba/telegram/callback"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type callbackHandler struct {
	bot         arumba.IBot
	readHandler CallbackHandler
}

func NewCallbackHandler(app arumba.Arumba, bot arumba.IBot) callbackHandler {
	return callbackHandler{
		bot: bot,
		readHandler: callback.ReadHandler{
			Bot: bot,
			EpisodeSearcher: episode.Search{
				Repo: app.EpisodeRepo,
			},
		},
	}
}

func (handler callbackHandler) Handle(event *tgbotapi.CallbackQuery) {
	command, arg := handler.extractData(event.Data)
	switch command {
	case ReadCallback:
		handler.readHandler.Handle(event.Message.Chat.ID, arg)
	default:
		handler.bot.NotifyError(errors.New("command not found : " + command))
	}
}

func (handler callbackHandler) extractData(data string) (string, string) {
	base64, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		handler.bot.NotifyError(err)
	}
	arr := strings.Split(string(base64), "_")

	return arr[0], arr[1]
}
