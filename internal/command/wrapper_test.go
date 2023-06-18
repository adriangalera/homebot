package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var command1 = remoteCommand{
	Command: "/test1",
	Path:    "/test1",
	Port:    1,
}
var command2 = remoteCommand{
	Command: "/test2",
	Path:    "/test2",
	Port:    2,
}

func TestArrayListWrapper_FindByCommand(t *testing.T) {
	type fields struct {
		commands []Command
	}
	type args struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Command
	}{
		{name: "Finds first command", fields: fields{commands: []Command{command1, command2}}, args: args{text: "/test1"}, want: command1},
		{name: "Finds second command", fields: fields{commands: []Command{command1, command2}}, args: args{text: "/test2"}, want: command2},
		{name: "Don't find command", fields: fields{commands: []Command{command1, command2}}, args: args{text: "/xxx"}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayListWrapper{
				commands: tt.fields.commands,
			}
			assert.Equalf(t, tt.want, a.FindByCommand(tt.args.text), "FindByCommand(%v)", tt.args.text)
		})
	}
}

func TestArrayListWrapper_GetAll(t *testing.T) {
	type fields struct {
		commands []Command
	}
	tests := []struct {
		name   string
		fields fields
		want   []Command
	}{
		{name: "Returns all commands", fields: fields{commands: []Command{command1, command2}}, want: []Command{command1, command2}},
		{name: "Returns empty commands", fields: fields{commands: []Command{}}, want: []Command{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayListWrapper{
				commands: tt.fields.commands,
			}
			assert.Equalf(t, tt.want, a.GetAll(), "GetAll()")
		})
	}
}

func TestArrayListWrapper_Length(t *testing.T) {
	type fields struct {
		commands []Command
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "Two commands", fields: fields{commands: []Command{command1, command2}}, want: 2},
		{name: "No commands", fields: fields{commands: []Command{}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayListWrapper{
				commands: tt.fields.commands,
			}
			assert.Equalf(t, tt.want, a.Length(), "Length()")
		})
	}
}
