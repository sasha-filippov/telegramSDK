package main

import (
	"fmt"
	"log"
	"telegramSDK/telegramClient"
)

func main() {
	token := "5778811258:AAHw38nJ9hVVvnh1HslfuohwCQfic2UjOZ0"

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
