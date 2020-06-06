package main

import (
	"bytes"
	"fmt"
	"image/png"
	"log"

	dc "github.com/NineStems/dishonoredCrypt"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

const (
	BotToken = "ur token"
)

func main() {
	// подключаемся к боту с помощью токена
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	upd, _ := bot.GetUpdatesChan(ucfg)
	for {
		select {
		case update := <-upd:
			ChatID := update.Message.Chat.ID

			img, err := dc.GetImageByWords(update.Message.Text)

			if err != nil {
				fmt.Println(err)
			}
			buff := new(bytes.Buffer)
			png.Encode(buff, img)
			if err != nil {
				fmt.Println("failed to create buffer", err)
			}
			b := buff.Bytes()

			fb := tgbotapi.FileBytes{"test", b}
			if err != nil {
				log.Println(err)
				continue
			}
			t := tgbotapi.NewPhotoUpload(ChatID, fb)
			bot.Send(t)
		}
	}
}
