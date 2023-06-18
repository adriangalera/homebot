package command

import (
	"errors"
	"gopkg.in/yaml.v3"
	"homebot/internal/io"
	"os"
)

func readAllCommandsFile(path string) ([]remoteCommand, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return []remoteCommand{}, err
	}
	var commands []remoteCommand
	for _, e := range entries {
		filename := path + string(os.PathSeparator) + e.Name()
		fileContent, errorReadingFile := os.ReadFile(filename)
		var cmd remoteCommand
		errUnMarshalling := yaml.Unmarshal(fileContent, &cmd)
		if errorReadingFile != nil || errUnMarshalling != nil {
			return []remoteCommand{}, errors.New("cannot read file " + filename)
		}
		commands = append(commands, cmd)
	}
	return commands, nil
}

func AddLocalCommands(wrapper Wrapper) {
	helpCommand := HelpCommand{wrapper: wrapper}
	wrapper.Append(&helpCommand)
}

func ReadCommands(commandFolder string) (Wrapper, error) {
	commandFolderExists, err := io.ExistsFile(commandFolder)
	if err != nil {
		return nil, err
	}
	if !commandFolderExists {
		err = io.CreateFolder(commandFolder)
		return &ArrayListWrapper{commands: []Command{}}, err
	}
	remoteCommands, errorReadingCommandFiles := readAllCommandsFile(commandFolder)
	if errorReadingCommandFiles != nil {
		return nil, errorReadingCommandFiles
	}
	var comands []Command
	for _, remoteCommand := range remoteCommands {
		comands = append(comands, remoteCommand)
	}
	wrapper := &ArrayListWrapper{commands: comands}
	return wrapper, nil
}
