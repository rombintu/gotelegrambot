package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rombintu/gotelegrambot/tools"
)

const (
	// APIEndpoint is the endpoint for all API methods,
	// with formatting for Sprintf.
	APIEndpoint = "https://api.telegram.org/bot%s/%s"
	// FileEndpoint is the endpoint for downloading a file from Telegram.
	// FileEndpoint = "https://api.telegram.org/file/bot%s/%s"
)

var botKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("button1"),
		tgbotapi.NewKeyboardButton("button2"),
	),
)

func main() {
	conf, err := tools.ParseConfigToml("config.toml")
	if err != nil {
		log.Fatalf("%v", err)
	}

	bot, err := tgbotapi.NewBotAPI(conf.Default.Token)
	if err != nil {
		log.Fatalf("TOKEN ERROR: %v", err)
	}

	bot.Debug = conf.Default.Debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = conf.Default.TimeoutUpdate

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("%v", err)
	}

	for u := range updates {

		if u.Message == nil {
			continue
		}

		userText := u.Message.Text
		userID := u.Message.Chat.ID
		// userLang := u.Message.From.LanguageCode
		// userName := u.Message.From.UserName

		printBot := func(text string) {
			message := tgbotapi.NewMessage(userID, text)
			if _, err := bot.Send(message); err != nil {
				log.Printf("ERROR SEND: %v", err)
			}
		}

		returnKeyboard := func(text string) {
			message := tgbotapi.NewMessage(userID, text)
			message.ReplyMarkup = botKeyboard
			if _, err := bot.Send(message); err != nil {
				log.Printf("ERROR RETURN KEYBOARD: %v", err)
			}
		}

		switch userText {
		case "/start":
			returnKeyboard("hello!")
		case "button1":
			printBot("button1")
		case "button2":
			printBot("button2")
		default:
			printBot("Неизвестная команда")
		}
	}
}
