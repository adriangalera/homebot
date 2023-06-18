package command

type Wrapper interface {
	FindByCommand(text string) Command
	Length() int
	GetAll() []Command
	Append(command Command)
}

type ArrayListWrapper struct {
	commands []Command
}

func (a *ArrayListWrapper) FindByCommand(text string) Command {
	for _, cmd := range a.commands {
		if cmd.GetCommand() == text {
			return cmd
		}
	}
	return nil
}

func (a *ArrayListWrapper) Length() int {
	return len(a.commands)
}

func (a *ArrayListWrapper) GetAll() []Command {
	return a.commands
}

func (a *ArrayListWrapper) Append(command Command) {
	a.commands = append(a.commands, command)
}
