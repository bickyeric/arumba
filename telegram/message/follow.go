package message

import (
	"github.com/bickyeric/arumba"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

type follow struct {
	bot arumba.IBot
}

// NewFollow ...
func NewFollow(bot arumba.Bot) Handler {
	return follow{bot}
}

func (f follow) Handle(message *tgbotapi.Message) {
	comicName := message.CommandArguments()

	if comicName == "" {
		f.showFollowed(message.Chat.ID)
	} else {
		f.follow(message.Chat.ID, comicName)
	}
}

func (f follow) showFollowed(chatID int64) {
	log.Println("showing followed comic")
}

func (f follow) follow(chatID int64, comicName string) {
	log.Println("following new comic")
}
