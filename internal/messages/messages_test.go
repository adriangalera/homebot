package messages

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"homebot/internal/config"
	"testing"
)

func TestSendReadyMessage(t *testing.T) {
	testConfiguration := config.Config{
		Bot: config.BotConfig{Name: "TestSendReadyMessage"}}

	okBot := &MockHomeBot{}
	okBot.On("SendText", "TestSendReadyMessage ready!").Return(nil)
	failingBot := &MockHomeBot{}
	failingBot.On("SendText", mock.Anything).Return(errors.New("failing sender"))

	okMessages := NewMessages(&testConfiguration, okBot)
	failingMessage := NewMessages(&testConfiguration, failingBot)

	tests := []struct {
		name     string
		messages *Messages
		err      error
	}{
		{name: "Test sending correctly", messages: okMessages, err: nil},
		{name: "Test sending fails", messages: failingMessage, err: errors.New("failing sender")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.messages.SendReadyMessage()
			assert.Equal(t, err, tt.err)
		})
	}
}
