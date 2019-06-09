package message

import "github.com/go-telegram-bot-api/telegram-bot-api"

// ...
var (
	FeedbackCommand = "feedback"
	FollowCommand   = "follow"
	HelpCommand     = "help"
	ReadCommand     = "read"
	StartCommand    = "start"
	GenericCommand  = "generic"
)

// Handler ...
type Handler interface {
	Handle(message *tgbotapi.Message)
}
