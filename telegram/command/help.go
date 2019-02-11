package command

import (
	"github.com/bickyeric/arumba/telegram"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Help struct {
	Bot telegram.Bot
}

func (h Help) Handle(message *tgbotapi.Message) {
	helpMsg := tgbotapi.NewMessage(message.Chat.ID, "Join channel t.me/nbcomic, untuk selalu update comic terbaru dari berbagai sumber :D\n\n klik /feedback untuk ngasih feedback atau masukan ke developer :D")
	h.Bot.Send(helpMsg)
}
