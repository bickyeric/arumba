package message

import (
	"github.com/bickyeric/arumba"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type help struct {
	bot arumba.IBot
}

// NewHelp ...
func NewHelp(bot arumba.Bot) Handler {
	return help{bot}
}

func (h help) Handle(message *tgbotapi.Message) {
	h.bot.SendTextMessage(message.Chat.ID, "Join channel t.me/arumba_channel, untuk selalu update comic terbaru dari berbagai sumber :D\n\nGunakan /feedback untuk ngasih feedback atau masukan ke developer :D")
}
