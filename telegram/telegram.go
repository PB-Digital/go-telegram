package telegram

import (
	"errors"
)

const (
	defaultBaseURL = "https://api.telegram.org"
)

func NewClient(token string) (*Client, error) {
	if token == "" {
		err := errors.New("unauthorized: no token present")
		return nil, err
	}
	return &Client{token}, nil
}
