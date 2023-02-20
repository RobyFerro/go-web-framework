package registers

type command interface {
	Run()
	Register()
}

// CommandRegister defines all registered commands
type CommandRegister map[string]command

// Add new command to cli register
func (c CommandRegister) Add(key string, value command) {
	c[key] = value
}
