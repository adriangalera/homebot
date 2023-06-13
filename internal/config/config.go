package config

import (
	"errors"
	"github.com/spf13/viper"
	"homebot/internal/io"
	"log"
)

type Config struct {
	Bot      BotConfig      `mapstructure:"bot"`
	Telegram TelegramConfig `mapstructure:"telegram"`
	Api      ApiConfig      `mapstructure:"api"`
}

type BotConfig struct {
	Name string `mapstructure:"name"`
}

type TelegramConfig struct {
	Token   string `mapstructure:"token"`
	GroupId string `mapstructure:"group-id"`
}

type ApiConfig struct {
	Port int `mapstructure:"port"`
}

func validate(newConfig *Config) error {
	emptyBot := BotConfig{}
	if newConfig.Bot == emptyBot {
		return errors.New("bot not defined")
	}

	if newConfig.Telegram.Token == "" || newConfig.Telegram.GroupId == "" {
		return errors.New("telegram not defined")
	}

	emptyApi := ApiConfig{}
	if newConfig.Api == emptyApi {
		return errors.New("api not defined")
	}

	return nil
}

func ReadConfigurationFromFile(filename string) (*Config, error) {
	viper.SetConfigFile(filename)
	newConfig := Config{}
	exits, err := io.ExistsFile(filename)
	if !exits || err != nil {
		log.Printf("Config file %s not found ...", filename)
		log.Printf("Creating an empty config file at %s", filename)
		errorWritingEmptyConfigFile := io.WriteYamlFile(filename, newConfig)
		if errorWritingEmptyConfigFile != nil {
			return nil, errorWritingEmptyConfigFile
		}
		return nil, errors.New("empty configuration, file created")
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	_ = viper.Unmarshal(&newConfig)

	errorsValidating := validate(&newConfig)
	if errorsValidating != nil {
		return nil, errorsValidating
	}

	return &newConfig, nil
}
