package telegram

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

type Message struct {
	ChatId    string   `json:"chatId"`
	Body      string   `json:"body"`
	PhotoUrl  string   `json:"photoUrl"`
	ParseMode string   `json:"parseMode"`
	Document  Document `json:"document"`
}

type Document struct {
	Content []byte `json:"content"`
	Name    string `json:"name"`
	Caption string `json:"caption"`
}

type FormDataItem struct {
	Key      string
	Filename string
	Content  []byte
}

type Client struct {
	botToken string
}

func (client *Client) SendMessage(message Message) {
	endpoint := buildEndpoint(client.botToken, message)
	if hasDocument(message) {
		sendDocument(endpoint, message)
	} else {
		httpClient := &http.Client{Timeout: 30 * time.Second}
		httpClient.Get(endpoint)
	}
}

func sendDocument(endpoint string, message Message) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	bodyWriter.WriteField("caption", message.Document.Caption)
	bodyWriter.WriteField("chat_id", message.ChatId)

	var fileWriter io.Writer

	fileWriter, _ = bodyWriter.CreateFormFile("document", message.Document.Name)
	fileWriter.Write(message.Document.Content)

	contentType := fmt.Sprintf("multipart/form-data; boundary=%s", bodyWriter.Boundary())
	bodyWriter.Close()

	httpClient := &http.Client{Timeout: 30 * time.Second}
	httpClient.Post(endpoint, contentType, bodyBuf)
}

func buildEndpoint(botToken string, message Message) string {
	endpoint := fmt.Sprintf("%s/bot%s", defaultBaseURL, botToken)
	if hasDocument(message) {
		endpoint += "/sendDocument"
	} else if hasPhoto(message) {
		endpoint += fmt.Sprintf("/sendPhoto?chat_id=%s&photo=%s&caption=%s&parse_mode=%s",
			message.ChatId, message.PhotoUrl, message.Body, message.ParseMode)
	} else {
		endpoint += fmt.Sprintf("/sendMessage?chat_id=%s&text=%s&parse_mode=%s",
			message.ChatId, message.Body, message.ParseMode)
	}
	return endpoint
}

func hasDocument(message Message) bool {
	return len(message.Document.Name) > 0 && len(message.Document.Content) > 0
}

func hasPhoto(message Message) bool {
	return message.PhotoUrl != ""
}
