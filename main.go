package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// Episode merepresentasikan objek episode
type Episode struct {
	Page []string `json:"page"`
}

type Source struct {
	Name string `json:"name"`
}

// Comic merepresentasikan objek komik
type Comic struct {
	Source  Source  `json:"source"`
	Episode Episode `json:"episode"`
}

// PhotoParams ...
type PhotoParams struct {
	ChatID int64  `json:"chat_id"`
	Photo  string `json:"photo"`
}

func handleStartCommand(message *tgbotapi.Message) {
	arg := message.CommandArguments()
	if arg == "" {
		return
	}

	decodedArg, err := base64.StdEncoding.DecodeString(arg)
	if err != nil {
		log.Fatal(err)
	}
	decodedString := string(decodedArg)
	splittedString := strings.Split(decodedString, "_")

	resp, _ := http.Get("https://backend-bot.000webhostapp.com/index.php/comic/read/userid/" + splittedString[0] + "/" + splittedString[1])
	jsonRaw, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	dec := json.NewDecoder(strings.NewReader(string(jsonRaw)))
	var comic Comic
	if err := dec.Decode(&comic); err != nil {
		log.Fatal(err)
	}
	log.Print(comic.Source.Name)
	url := "https://api.telegram.org/bot" + os.Getenv("telegramToken") + "/sendPhoto"

	for _, page := range comic.Episode.Page {
		params := PhotoParams{message.Chat.ID, page}
		jsonParams, _ := json.Marshal(params)

		resp, err = http.Post(url, "application/json", bytes.NewBuffer(jsonParams))
		log.Print(resp.Body)
	}
}

var bot tgbotapi.BotAPI

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("telegramToken"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe("0.0.0.0:"+port, nil)

	for update := range updates {
		command := update.Message.Command()
		if command == "start" {
			handleStartCommand(update.Message)
			tqMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Spesial Thanks to : Mangacanblog")
			bot.Send(tqMsg)
		}
	}
}
