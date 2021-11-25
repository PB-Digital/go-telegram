package main

import (
	"fmt"
	"go-telegram/telegram"
	"io/ioutil"
)

func main() {
	var client, _ = telegram.NewClient("TELEGRAM_BOT_TOKEN")
	fileName := "FILE_NAME"
	fileContent, _ := ioutil.ReadFile(fmt.Sprintf("%s/%s", "FILE_BASE_URL", fileName))
	message := telegram.Message{
		ChatId: "TELEGRAM_GROUP_OR_CHANNEL_ID",
		Body:   "MESSAGE",
		Document: telegram.Document(struct {
			Content []byte
			Name    string
			Caption string
		}{Content: fileContent, Name: fileName}),
	}

	client.SendMessage(message)
}
