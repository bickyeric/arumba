package callback

import (
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// SelectComicCallback ...
var SelectComicCallback = "select-comic"

// SelectEpisodeCallback ...
var SelectEpisodeCallback = "select-episode"

// Handler ...
type Handler interface {
	Handle(event *tgbotapi.CallbackQuery)
}

// ExtractData ...
func ExtractData(data string) (string, string) {
	arr := strings.Split(data, "_")

	return arr[0], data[len(arr[0])+1:]
}
