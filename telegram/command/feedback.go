package command

import (
	"github.com/bickyeric/arumba"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var feedbackRequest = "Masukan kamu sangat berarti buat kami :D"

type FeedbackHandler struct {
	Bot arumba.IBot
}

func (f FeedbackHandler) Handle(message *tgbotapi.Message) {
	f.Bot.SendReplyMessage(message.Chat.ID, feedbackRequest)
}
