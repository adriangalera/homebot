package command

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"homebot/internal/io"
	"testing"
)

func TestReadCommands(t *testing.T) {
	commandFolder := "./testCommands"
	commands, err := ReadCommands(commandFolder)
	expectedCommands := []Command{
		remoteCommand{
			Command:     "/test",
			Path:        "/test",
			Port:        1234,
			Description: "test",
		},
		remoteCommand{
			Command:     "/test2",
			Path:        "/test2",
			Port:        1234,
			Description: "test2",
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expectedCommands, commands.GetAll())
}

func TestReadCommandsNotExistingFolder(t *testing.T) {
	randomId := uuid.New().String()
	commandFolder := "/tmp/" + randomId
	commands, err := ReadCommands(commandFolder)
	assert.Nil(t, err)
	assert.Equal(t, 0, commands.Length())
	existsFolder, err := io.ExistsFile(commandFolder)
	assert.True(t, existsFolder)
}

func TestAddLocalCommands(t *testing.T) {
	wrapper := &ArrayListWrapper{}
	type args struct {
		wrapper Wrapper
	}
	type expected struct {
		localCommands []Command
	}
	tests := []struct {
		name string
		args args
		expected
	}{
		{
			name:     "Add local commands",
			args:     args{wrapper: wrapper},
			expected: expected{localCommands: []Command{&HelpCommand{wrapper}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddLocalCommands(tt.args.wrapper)
			assert.Equal(t, tt.expected.localCommands, wrapper.GetAll())
		})
	}
}
