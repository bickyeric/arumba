package command

import (
	"os"
	"strconv"

	"github.com/bickyeric/arumba/service/comic"
	"github.com/bickyeric/arumba/telegram"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type common struct {
	bot           telegram.IBot
	comicSearcher comic.Search
}

func Common(bot telegram.IBot, comicSearcher comic.Search) telegram.CommandHandler {
	return common{bot, comicSearcher}
}

func (c common) Handle(message *tgbotapi.Message) {
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

func (c common) handleReadComic(message *tgbotapi.Message) {
	comics, err := c.comicSearcher.Perform(message.Text)
	if err != nil {
		c.bot.NotifyError(err)
	}

	c.bot.SendComicSelector(message.Chat.ID, comics)
}

func (c common) handleFeedback(message *tgbotapi.Message) {
	replyMessage := tgbotapi.NewMessage(message.Chat.ID, "Makasih masukannya...")
	replyMessage.ReplyToMessageID = message.MessageID
	c.bot.Bot().Send(replyMessage)

	chatID, _ := strconv.Atoi(os.Getenv("CHAT_ID"))
	forwardFeedbackMessage := tgbotapi.NewForward(int64(chatID), message.Chat.ID, message.MessageID)
	c.bot.Bot().Send(forwardFeedbackMessage)
}
