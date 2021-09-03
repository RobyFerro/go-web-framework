package register

// CommandRegister defines all registered commands
type CommandRegister map[string]interface{}

// Command is used to define a CLI command
type Command struct {
	Signature   string
	Description string
	Args        string
}
