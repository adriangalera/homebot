package main

import (
	"homebot/internal/api"
	"homebot/internal/bot"
	"homebot/internal/command"
	"homebot/internal/config"
	"homebot/internal/messages"
	"homebot/internal/telegram"
	"log"
	"net/http"
	"strconv"
)

func apiListen(mainConfig *config.Config, sender bot.HomeBot) {
	api.RegisterApiEndpoints(sender)
	apiAddr := "localhost:" + strconv.Itoa(mainConfig.Api.Port)
	log.Printf("API will listen at %s\n", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr, nil))
}

func sendReadyMessage(mainConfig *config.Config, sender bot.HomeBot) error {
	msg := messages.NewMessages(mainConfig, sender)
	return msg.SendReadyMessage()
}

func initializeTelegram(mainConfig *config.Config, wrapper command.Wrapper) (bot.HomeBot, error) {
	sender, errorStartTelegram := telegram.NewTelegramBot(mainConfig, &command.Executor{}, wrapper)
	if errorStartTelegram != nil {
		return &telegram.Bot{}, errorStartTelegram
	}
	return &sender, nil
}

func readCommands() command.Wrapper {
	commands, err := command.ReadCommands("/etc/homebot/commands")
	command.AddLocalCommands(commands)
	if err != nil {
		log.Panicf("Error: %s", err)
	}
	log.Printf("Available commands:")
	for _, cmd := range commands.GetAll() {
		log.Printf("%s", cmd.GetCommand())
	}
	return commands
}

func readConfiguration() (*config.Config, error) {
	mainConfig, errorReadingConfig := config.ReadConfigurationFromFile("/etc/homebot/config.yml")
	if errorReadingConfig != nil {
		return nil, errorReadingConfig
	}
	return mainConfig, nil
}

func main() {
	mainConfig, readConfigError := readConfiguration()
	if readConfigError != nil {
		log.Fatalf("Cannot read the configuration Error: %s", readConfigError)
		return
	}

	wrapper := readCommands()

	homeBot, telegramError := initializeTelegram(mainConfig, wrapper)
	if telegramError != nil {
		log.Fatalf("Cannot initialize messages. Error: %s", telegramError)
		return
	}

	errorSendingMessage := sendReadyMessage(mainConfig, homeBot)
	if errorSendingMessage != nil {
		log.Fatalf("Error sending message: %v", errorSendingMessage)
		return
	}
	go func() {
		apiListen(mainConfig, homeBot)
	}()
	homeBot.ListenForCommands()
}
