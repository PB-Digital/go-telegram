package main

import (
	"go-telegram/telegram"
)

func main() {
	var client, _ = telegram.NewClient("TELEGRAM_BOT_TOKEN")

	message := telegram.Message{
		ChatId: "TELEGRAM_GROUP_OR_CHANNEL_ID",
		Body:   "MESSAGE",
	}

	client.SendMessage(message)
}
