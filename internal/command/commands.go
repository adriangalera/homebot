package command

type Command interface {
	GetCommand() string
	GetDescription() string
}

type remoteCommand struct {
	Command     string `mapstructure:"command"`
	Description string `mapstructure:"description"`
	Path        string `mapstructure:"path"` //for simplicity all the request will be POST to localhost
	Port        int    `mapstructure:"port"`
}

type localCommand interface {
	Execute() string
}

func (r remoteCommand) GetCommand() string {
	return r.Command
}

func (r remoteCommand) GetDescription() string {
	return r.Description
}
