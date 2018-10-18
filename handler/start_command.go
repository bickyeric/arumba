package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/model"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// PhotoParams ...
type PhotoParams struct {
	ChatID int64  `json:"chat_id"`
	Photo  string `json:"photo"`
}

func parseArg(arg string) (string, int) {
	decodedArg, err := base64.StdEncoding.DecodeString(arg)
	if err != nil {
		log.Fatal(err)
	}

	decodedString := string(decodedArg)
	splittedString := strings.Split(decodedString, "_")
	episode, _ := strconv.Atoi(splittedString[1])

	return splittedString[0], episode
}

func StartCommand(message *tgbotapi.Message) {
	arg := message.CommandArguments()
	if arg == "" {
		bot := arumba.Instance()
		tqMsg := tgbotapi.NewMessage(message.Chat.ID, "hai")
		bot.Send(tqMsg)
		return
	}

	comicName, episode := parseArg(arg)
	comic := model.ReadComic(comicName, episode, 0)
	log.Print(comic.Source.Name)
	url := "https://api.telegram.org/bot" + os.Getenv("telegramToken") + "/sendPhoto"

	for _, page := range comic.Episode.Page {
		params := PhotoParams{message.Chat.ID, page}
		jsonParams, _ := json.Marshal(params)

		resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonParams))
		log.Print(resp.Body)
	}
}
