package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

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
	json, _ := ioutil.ReadAll(resp.Body)
	log.Printf("%s\n", json)
	resp.Body.Close()
}

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
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Masih Development mang!!!")
		bot.Send(msg)
	}
}
