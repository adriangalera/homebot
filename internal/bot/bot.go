package bot

import (
	"homebot/internal/command"
)

//go:generate mockery --name HomeBot
type HomeBot interface {
	SendText(message string) error
	SendImage(filename string) error
	SendAudio(filename string) error
	ListenForCommands() error
	HandleCommand(command command.Command) error
}
