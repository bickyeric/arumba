package message

import (
	"github.com/bickyeric/arumba"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

// FollowHandler ...
type FollowHandler struct {
	Bot arumba.IBot
}

// Handle ...
func (f FollowHandler) Handle(message *tgbotapi.Message) {
	comicName := message.CommandArguments()

	if comicName == "" {
		f.showFollowed(message.Chat.ID)
	} else {
		f.follow(message.Chat.ID, comicName)
	}
}

func (f FollowHandler) showFollowed(chatID int64) {
	log.Println("showing followed comic")
}

func (f FollowHandler) follow(chatID int64, comicName string) {
	log.Println("following new comic")
}
