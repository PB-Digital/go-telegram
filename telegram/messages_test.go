package telegram

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var botToken = "TOKEN"

var message = Message{
	ChatId:    "-123456",
	Body:      "Hello there!",
	ParseMode: "text",
}

var client, _ = NewClient(botToken)

func TestSendingPlainMessage(t *testing.T) {
	actualEndpoint := buildEndpoint(botToken, message)
	expectedEndpoint := "https://api.telegram.org/botTOKEN/sendMessage?chat_id=-123456&text=Hello there!&parse_mode=text"
	client.SendMessage(message)
	assert.Equal(t, expectedEndpoint, actualEndpoint)
}

func TestSendingHtmlMessage(t *testing.T) {
	htmlMessage := message
	htmlMessage.ParseMode = "html"
	actualEndpoint := buildEndpoint(botToken, htmlMessage)
	expectedEndpoint := "https://api.telegram.org/botTOKEN/sendMessage?chat_id=-123456&text=Hello there!&parse_mode=html"
	client.SendMessage(message)
	assert.Equal(t, expectedEndpoint, actualEndpoint)
}

func TestSendingWithPhoto(t *testing.T) {
	message.PhotoUrl = "https://test.test/test-image.png"
	actualEndpoint := buildEndpoint(botToken, message)
	expectedEndpoint := "https://api.telegram.org/botTOKEN/sendPhoto?chat_id=-123456&photo=https://test.test/test-image.png&caption=Hello there!&parse_mode=text"
	client.SendMessage(message)
	assert.Equal(t, expectedEndpoint, actualEndpoint)
}

func TestSendingWithDocument(t *testing.T) {
	message.Document = Document(struct {
		Content []byte
		Name    string
		Caption string
	}{Content: []byte{'a', 'b', 'c'}, Name: "test.txt", Caption: "Test file"})
	actualEndpoint := buildEndpoint(botToken, message)
	expectedEndpoint := "https://api.telegram.org/botTOKEN/sendDocument"
	client.SendMessage(message)
	sendDocument(actualEndpoint, message)
	assert.Equal(t, expectedEndpoint, actualEndpoint)
}
