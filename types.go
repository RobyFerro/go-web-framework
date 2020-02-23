package main

// CommandRegister defines all registered commands
type CommandRegister = map[string]interface{}

// ControllerRegister defines a controller register type.
// This will be used to resolve this register in service container
type ControllerRegister = []interface{}

// ModelRegister defines a controller register type.
// This will be used to resolve this register in service container
type ModelRegister = []interface{}

// ServiceRegister defines a controller register type.
// This will be used to resolve this register in service container
type ServiceRegister = []interface{}
