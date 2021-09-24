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

	// _, err = bot.SetWebhook(tgbotapi.NewWebhook("https://webhook.site/14a926fb-364d-46ee-b01c-15e69130f41b/" + bot.Token))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// updates := bot.ListenForWebhook("/" + bot.Token)
	// go http.ListenAndServe("0.0.0.0:5000", nil)

	// for update := range updates {
	// 	log.Printf("%+v\n", update)
	// }

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("%v", err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		userText := update.Message.Text
		userCommand := update.Message.Command()
		userID := update.Message.Chat.ID

		// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(userID, "")
			switch userCommand {
			case "start":
				msg.ReplyMarkup = botKeyboard
				msg.Text = "Hello"
			case "help":
				msg.Text = "type /sayhi or /status."
			case "sayhi":
				msg.Text = "Hi :)"
			case "status":
				msg.Text = "I'm ok."
			case "withArgument":
				msg.Text = "You supplied the following argument: " + update.Message.CommandArguments()
			case "html":
				msg.ParseMode = "html"
				msg.Text = "This will be interpreted as HTML, click <a href=\"https://www.example.com\">here</a>"
			default:
				msg.Text = "I don't know that command"
			}
			bot.Send(msg)
		} else if update.Message.Document != nil {
			doc := update.Message.Document
			userFileID := doc.FileID
			// userFileName := doc.FileName
			msg := tgbotapi.NewDocumentShare(userID, userFileID)
			bot.Send(msg)
		} else {
			msg := tgbotapi.NewMessage(userID, "")
			switch userText {
			case "button1":
				msg.Text = "press button1"
			case "button2":
				photoBytes, photoName := tools.ReadFileFromUploads()
				newFileConfig := tgbotapi.FileBytes{
					Name:  photoName,
					Bytes: photoBytes,
				}
				msg.Text = photoName
				bot.Send(tgbotapi.NewPhotoUpload(userID, newFileConfig))
			default:
				msg.Text = "I don't know that command"
			}
			bot.Send(msg)
		}
	}

}
