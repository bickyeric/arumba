package command

import (
	"github.com/bickyeric/arumba/telegram"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type help struct {
	bot telegram.IBot
}

func Help(bot telegram.IBot) telegram.CommandHandler {
	return help{bot}
}

func (h help) Handle(message *tgbotapi.Message) {
	h.bot.SendTextMessage(message.Chat.ID, "Join channel t.me/nbcomic, untuk selalu update comic terbaru dari berbagai sumber :D\n\n klik /feedback untuk ngasih feedback atau masukan ke developer :D")
}
