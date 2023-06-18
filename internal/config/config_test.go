package config

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_readConfigurationFromFile(t *testing.T) {
	expectedConfig := &Config{
		Bot: BotConfig{Name: "bot-name"},
		Telegram: TelegramConfig{
			Token:   "token",
			GroupId: "group-id",
		},
		Api: ApiConfig{Port: 1234},
	}

	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{name: "Test read configuration file", filename: "./examples/config.yml", wantErr: false},
		{name: "Test read api not defined", filename: "examples/config-no-api.yml", wantErr: true},
		{name: "Test read bot not defined", filename: "examples/config-no-bot.yml", wantErr: true},
		{name: "Test read telegram not defined", filename: "examples/config-no-telegram.yml", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readConfig, err := ReadConfigurationFromFile(tt.filename)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}
			assert.Equal(t, expectedConfig, readConfig)
		})
	}
}

func Test_readConfigurationFromFile_creates_empty(t *testing.T) {
	configFile := randomConfigFileName()
	assert.NoFileExists(t, configFile)
	_, _ = ReadConfigurationFromFile(configFile)
	assert.FileExists(t, configFile)
}

func randomConfigFileName() string {
	randomId := uuid.New().String()
	configFile := fmt.Sprintf("/tmp/test-config-%s/test-config.yml", randomId)
	return configFile
}
