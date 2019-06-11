package message

import (
	"os"
	"strconv"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/service/comic"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

type generic struct {
	bot           arumba.IBot
	notifier      arumba.BotNotifier
	comicSearcher comic.Searcher
}

// NewGeneric ...
func NewGeneric(bot arumba.Bot,
	notifier arumba.BotNotifier,
	comicSearcher comic.Searcher) Handler {
	return generic{
		bot:           bot,
		notifier:      bot,
		comicSearcher: comicSearcher,
	}
}

func (c generic) Handle(message *tgbotapi.Message) {
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

func (c generic) handleReadComic(message *tgbotapi.Message) {
	contextLog := log.WithFields(
		log.Fields{
			"chat_id": message.Chat.ID,
		},
	)
	comics, err := c.comicSearcher.Perform(message.Text)
	if err != nil {
		c.notifier.NotifyError(err)
	}

	if len(comics) < 1 {
		c.bot.SendNotFoundComic(message.Chat.ID, message.Text)
		contextLog.Info("Not found comic name sent")
	} else {
		c.bot.SendComicSelector(message.Chat.ID, comics)
		contextLog.Info("Comic selector sent")
	}
}

func (c generic) handleFeedback(message *tgbotapi.Message) {
	replyMessage := tgbotapi.NewMessage(message.Chat.ID, "Makasih masukannya...")
	replyMessage.ReplyToMessageID = message.MessageID
	c.bot.Bot().Send(replyMessage)

	chatID, _ := strconv.Atoi(os.Getenv("CHAT_ID"))
	forwardFeedbackMessage := tgbotapi.NewForward(int64(chatID), message.Chat.ID, message.MessageID)
	c.bot.Bot().Send(forwardFeedbackMessage)
}
