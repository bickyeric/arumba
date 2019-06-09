package message

import (
	"github.com/bickyeric/arumba"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var feedbackRequest = "Masukan kamu sangat berarti buat kami :D"

type feedback struct {
	bot arumba.IBot
}

// NewFeedback ...
func NewFeedback(bot arumba.Bot) Handler {
	return feedback{
		bot: bot,
	}
}

func (f feedback) Handle(message *tgbotapi.Message) {
	f.bot.SendReplyMessage(message.Chat.ID, feedbackRequest)
}
