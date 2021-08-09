package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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
		tgbotapi.NewKeyboardButton("hi"),
		tgbotapi.NewKeyboardButton("by"),
	),
)

func main() {
	var VALUE int

	valInit := func() {
		nanoHelper := rand.NewSource(time.Now().UnixNano())
		VALUE = rand.New(nanoHelper).Intn(100)
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOTOKEN"))
	if err != nil {
		log.Fatalf("TOKEN ERROR: %v", err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("%v", err)
	}

	valInit()

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
			returnKeyboard("Отгадай загаданное число (от 1 до 100)")
			// printBot(fmt.Sprint(VALUE))
			continue
		case fmt.Sprint(VALUE):
			printBot("Победа. Меняю число")
			valInit()
			continue
		}
		// default:
		// 	returnKeyboard("Неизвестная команда")
		// }

		userVal, err := strconv.ParseInt(userText, 10, 64)
		if err == nil {
			if int64(VALUE) > userVal {
				printBot("Бери больше")
			} else if int64(VALUE) < userVal {
				printBot("Бери меньше")
			}
		} else {
			printBot("Ожидается число")
		}

	}
}
