package telegram

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenIsOk(t *testing.T) {
	client, err := NewClient("TOKEN")
	assert.NotNil(t, client)
	assert.Nil(t, err)
}

func TestTokenIsEmpty(t *testing.T) {
	client, err := NewClient("")
	assert.Nil(t, client)
	assert.Equal(t, "unauthorized: no token present", err.Error())
}
