package handler

import (
	"github.com/bickyeric/arumba"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var feedback = "Masukan kamu sangat berarti buat kami :D"

func FeedbackCommand(message *tgbotapi.Message) {
	bot := arumba.Instance()

	helpMsg := tgbotapi.NewMessage(message.Chat.ID, feedback)
	helpMsg.ReplyMarkup = tgbotapi.ForceReply{ForceReply: true}
	bot.Send(helpMsg)
}

func handleFeedback(message *tgbotapi.Message) {
	bot := arumba.Instance()
	replyMessage := tgbotapi.NewMessage(message.Chat.ID, "Makasih masukannya...")
	replyMessage.ReplyToMessageID = message.MessageID
	bot.Send(replyMessage)

	forwardFeedbackMessage := tgbotapi.NewForward(610339834, message.Chat.ID, message.MessageID)
	bot.Send(forwardFeedbackMessage)
}
