package command

import "github.com/go-telegram-bot-api/telegram-bot-api"

// ...
var (
	FeedbackCommand = "feedback"
	HelpCommand     = "help"
	ReadCommand     = "read"
	StartCommand    = "start"
)

// CommandHandler ...
type CommandHandler interface {
	Handle(message *tgbotapi.Message)
}
