package telegram

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"homebot/internal/command"
	"homebot/internal/config"
	"log"
	"os"
)

type Bot struct {
	ctx             context.Context
	bot             *bot.Bot
	config          *config.Config
	commandExecutor *command.Executor
	wrapper         command.Wrapper
}

func NewTelegramBot(config *config.Config, executor *command.Executor, wrapper command.Wrapper) (Bot, error) {
	b, err := bot.New(config.Telegram.Token)
	if err != nil {
		return Bot{}, err
	}

	telegramHomeBot := Bot{bot: b, config: config, commandExecutor: executor, ctx: context.TODO(), wrapper: wrapper}
	return telegramHomeBot, nil
}

func (t *Bot) SendText(message string) error {
	_, err := t.bot.SendMessage(context.TODO(), &bot.SendMessageParams{
		ChatID: t.config.Telegram.GroupId,
		Text:   message,
	})
	return err
}

func (t *Bot) SendImage(filename string) error {
	fileContents, errorReadingFile := os.ReadFile(filename)

	if errorReadingFile != nil {
		return errorReadingFile
	}

	_, errorSendPhoto := t.bot.SendPhoto(context.TODO(), &bot.SendPhotoParams{
		ChatID: t.config.Telegram.GroupId,
		Photo:  &models.InputFileUpload{Filename: filename, Data: bytes.NewReader(fileContents)},
	})
	return errorSendPhoto
}

func (t *Bot) SendAudio(filename string) error {
	fileContents, errorReadingFile := os.ReadFile(filename)

	if errorReadingFile != nil {
		return errorReadingFile
	}

	_, errorSendPhoto := t.bot.SendAudio(context.TODO(), &bot.SendAudioParams{
		ChatID: t.config.Telegram.GroupId,
		Audio:  &models.InputFileUpload{Filename: filename, Data: bytes.NewReader(fileContents)},
	})
	return errorSendPhoto
}

func (t *Bot) ListenForCommands() error {
	log.Printf("Listening for telegram bot commands ...")

	for _, cmd := range t.wrapper.GetAll() {
		t.bot.RegisterHandler(bot.HandlerTypeMessageText, cmd.GetCommand(), bot.MatchTypeExact, t.commandHandler)
	}

	t.bot.Start(t.ctx)
	return nil
}

func (t *Bot) HandleCommand(command command.Command) error {
	response, err := t.commandExecutor.Execute(command)
	if response != "" {
		errorSendingResponse := t.SendText(response)
		if err != nil {
			log.Printf("couldn't send telegarm message. error: %s", errorSendingResponse)
		}
	}
	return err
}

func (t *Bot) commandHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	var messageError error
	commandText := update.Message.Text
	log.Printf("Received %v command", commandText)
	cmd := t.wrapper.FindByCommand(commandText)
	if cmd == nil {
		commandNotFoundText := fmt.Sprintf("remoteCommand not recognized %s", commandText)
		log.Printf(commandNotFoundText + "\n")
		messageError = t.SendText(commandNotFoundText)
	} else {
		errorExecutingCommand := t.HandleCommand(cmd)
		if errorExecutingCommand != nil {
			errorMessage := fmt.Sprintf("Cannot execute command %s", commandText)
			log.Printf("%s", errorMessage)
			messageError = t.SendText(errorMessage)
		}
	}
	if messageError != nil {
		log.Printf("couldn't send telegarm message. error: %s", messageError)
	}
}
