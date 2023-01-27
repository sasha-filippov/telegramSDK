package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"telegramSDK/telegramClient"
)

func main() {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("env")
	vp.AddConfigPath(".")
	err := vp.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	token := vp.GetString("Token")

	tgClient := telegramClient.NewTelegramClient(token)
	offset := 0

	for {
		updates, err := tgClient.Updates(offset)
		if err != nil {
			log.Fatal(err)
		}
		for _, update := range updates {
			err := tgClient.SendMessage(update.Message.Chat.ChatID, update.Message.Text)
			if err != nil {
				fmt.Println(err.Error())
			}
			offset = update.UpdateID + 1

		}
	}
}
