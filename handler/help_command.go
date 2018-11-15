package handler

import (
	"github.com/bickyeric/arumba"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func HelpCommand(message *tgbotapi.Message) {
	bot := arumba.Instance()

	helpMsg := tgbotapi.NewMessage(message.Chat.ID, "Join channel t.me/nbcomic, untuk selalu update comic terbaru dari berbagai sumber :D\n\n klik /feedback untuk ngasih feedback atau masukan ke developer :D")
	bot.Send(helpMsg)
}
