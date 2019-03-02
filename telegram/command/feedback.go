package command

import (
	"github.com/bickyeric/arumba/telegram"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var feedbackRequest = "Masukan kamu sangat berarti buat kami :D"

type feedback struct {
	bot telegram.IBot
}

func Feedback(bot telegram.IBot) telegram.CommandHandler {
	return feedback{bot}
}

func (f feedback) Handle(message *tgbotapi.Message) {
	f.bot.SendReplyMessage(message.Chat.ID, feedbackRequest)
}
