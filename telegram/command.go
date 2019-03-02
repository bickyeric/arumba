package telegram

import "github.com/go-telegram-bot-api/telegram-bot-api"

// ...
var (
	FeedbackCommand = "feedback"
	HelpCommand     = "help"
	ReadCommand     = "read"
	StartCommand    = "start"
)

type CommandHandler interface {
	Handle(message *tgbotapi.Message)
}
