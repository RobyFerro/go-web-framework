package console

type @@TMP@@ struct {
	register.Command
}

// Command registration
func (c *@@TMP@@) Register() {
	c.Signature = "command:signature"               // Change command signature
	c.Description = "Execute database seeder"       // Change command description
}

// Command registration
func (c *@@TMP@@) Help() {
	// Implement with command help message
	log.Println("Usage: ---------")
}

// Command business logic
func (c *@@TMP@@) Run() {
	// Insert command logic
}


