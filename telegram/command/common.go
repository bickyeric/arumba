package command

import (
	"os"
	"strconv"

	"github.com/bickyeric/arumba/service/comic"
	"github.com/bickyeric/arumba/telegram"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Common struct {
	Bot           telegram.Bot
	ComicSearcher comic.Search
}

func (c Common) Handle(message *tgbotapi.Message) {
	if message.ReplyToMessage != nil {
		switch message.ReplyToMessage.Text {
		case feedbackRequest:
			c.handleFeedback(message)
			return
		case comicNameRequest:
			c.handleReadComic(message)
		}
	}
}

func (c Common) handleReadComic(message *tgbotapi.Message) {
	comics, err := c.ComicSearcher.Perform(message.Text)
	if err != nil {
		c.Bot.NotifyError(err)
	}

	c.Bot.SendComicSelector(message.Chat.ID, comics)
}

func (c Common) handleFeedback(message *tgbotapi.Message) {
	replyMessage := tgbotapi.NewMessage(message.Chat.ID, "Makasih masukannya...")
	replyMessage.ReplyToMessageID = message.MessageID
	c.Bot.Send(replyMessage)

	chatID, _ := strconv.Atoi(os.Getenv("CHAT_ID"))
	forwardFeedbackMessage := tgbotapi.NewForward(int64(chatID), message.Chat.ID, message.MessageID)
	c.Bot.Send(forwardFeedbackMessage)
}
