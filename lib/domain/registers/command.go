package registers

// CommandRegister defines all registered commands
type CommandRegister map[string]interface{}

// Add new command to cli register
func (c CommandRegister) Add(key string, value interface{}) {
	if len(c) == 0 {
		c = map[string]interface{}{}
	}

	c[key] = value
}
