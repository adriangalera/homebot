package command

type HelpCommand struct {
	wrapper Wrapper
}

func (h *HelpCommand) GetCommand() string {
	return "/help"
}

func (h *HelpCommand) GetDescription() string {
	return "shows all commands available"
}

func (h *HelpCommand) Execute() string {
	return h.commandsToString()
}

func (h *HelpCommand) commandsToString() string {
	helpMessage := "Available commands:\n\n"
	if h.wrapper != nil {
		for _, cmd := range h.wrapper.GetAll() {
			helpMessage += cmd.GetCommand() + "\t\t" + cmd.GetDescription() + "\n"
		}
	}
	return helpMessage
}
