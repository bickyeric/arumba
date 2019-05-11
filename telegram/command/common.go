package command

import (
	"os"
	"strconv"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/service/comic"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

type CommonHandler struct {
	Bot           arumba.IBot
	ComicSearcher comic.Search
}

func (c CommonHandler) Handle(message *tgbotapi.Message) {
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

func (c CommonHandler) handleReadComic(message *tgbotapi.Message) {
	contextLog := log.WithFields(
		log.Fields{
			"chat_id": message.Chat.ID,
		},
	)
	comics, err := c.ComicSearcher.Perform(message.Text)
	if err != nil {
		c.Bot.NotifyError(err)
	}

	if len(comics) < 1 {
		c.Bot.SendNotFoundComic(message.Chat.ID, message.Text)
		contextLog.Info("Not found comic name sent")
	} else {
		c.Bot.SendComicSelector(message.Chat.ID, comics)
		contextLog.Info("Comic selector sent")
	}
}

func (c CommonHandler) handleFeedback(message *tgbotapi.Message) {
	replyMessage := tgbotapi.NewMessage(message.Chat.ID, "Makasih masukannya...")
	replyMessage.ReplyToMessageID = message.MessageID
	c.Bot.Bot().Send(replyMessage)

	chatID, _ := strconv.Atoi(os.Getenv("CHAT_ID"))
	forwardFeedbackMessage := tgbotapi.NewForward(int64(chatID), message.Chat.ID, message.MessageID)
	c.Bot.Bot().Send(forwardFeedbackMessage)
}
