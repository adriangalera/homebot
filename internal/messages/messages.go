package messages

import (
	"fmt"
	"homebot/internal/bot"
	"homebot/internal/config"
)

type Messages struct {
	config *config.Config
	bot    bot.HomeBot
}

func NewMessages(config *config.Config, bot bot.HomeBot) *Messages {
	return &Messages{
		config: config,
		bot:    bot,
	}
}

func (m *Messages) SendReadyMessage() error {
	msg := fmt.Sprintf("%s ready!", m.config.Bot.Name)
	return m.bot.SendText(msg)
}
