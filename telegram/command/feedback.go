package command

import (
	"github.com/bickyeric/arumba/telegram"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var feedbackRequest = "Masukan kamu sangat berarti buat kami :D"

// Feedback ...
type Feedback struct {
	Bot telegram.Bot
}

// Handle ...
func (f Feedback) Handle(message *tgbotapi.Message) {
	f.Bot.SendReplyMessage(message.Chat.ID, feedbackRequest)
}
