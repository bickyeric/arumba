package arumba

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	cachedBot *tgbotapi.BotAPI
	botMutex  sync.Once
)

func NewBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("telegramToken"))
	if err != nil {
		log.Fatal(err)
	}

	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	bot.Debug = debug

	return bot
}

func Instance() *tgbotapi.BotAPI {
	botMutex.Do(func() {
		cachedBot = NewBot()
	})

	return cachedBot
}
