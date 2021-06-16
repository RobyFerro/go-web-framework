package register

// CommandRegister defines all registered commands
type CommandRegister struct {
	List map[string]interface{}
}

// ControllerRegister defines a controller register type.
// This will be used to resolve this register in service container
type ControllerRegister struct {
	List []interface{}
}

// ModelRegister defines a controller register type.
// This will be used to resolve this register in service container
type ModelRegister struct {
	List []interface{}
}

// ServiceRegister defines a controller register type.
// This will be used to resolve this register in service container
type ServiceRegister struct {
	List []interface{}
}
