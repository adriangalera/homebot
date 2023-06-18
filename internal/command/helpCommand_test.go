package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelpCommand_Execute(t *testing.T) {
	cmd := &HelpCommand{}
	cmd2 := remoteCommand{Command: "test", Description: "test"}
	emptyHelp := "Available commands:\n\n"
	helpMessage := "Available commands:\n\n/help\t\tshows all commands available\ntest\t\ttest\n"
	type fields struct {
		wrapper Wrapper
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "nil wrapper", fields: fields{wrapper: nil}, want: emptyHelp},
		{name: "empty wrapper", fields: fields{wrapper: &ArrayListWrapper{commands: []Command{}}}, want: emptyHelp},
		{name: "prints all commands formatted", fields: fields{wrapper: &ArrayListWrapper{commands: []Command{cmd, cmd2}}}, want: helpMessage},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			he := HelpCommand{
				wrapper: tt.fields.wrapper,
			}
			assert.Equalf(t, tt.want, he.Execute(), "Execute()")
		})
	}
}

func TestHelpCommand_GetCommand(t *testing.T) {
	help := &HelpCommand{}
	assert.Equal(t, "/help", help.GetCommand())
}
