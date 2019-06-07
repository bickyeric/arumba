package message

import (
	"github.com/bickyeric/arumba"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type HelpHandler struct {
	Bot arumba.IBot
}

func (h HelpHandler) Handle(message *tgbotapi.Message) {
	h.Bot.SendTextMessage(message.Chat.ID, "Join channel t.me/arumba_channel, untuk selalu update comic terbaru dari berbagai sumber :D\n\nGunakan /feedback untuk ngasih feedback atau masukan ke developer :D")
}
