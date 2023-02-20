package registers

// ControllerRegister defines a controller register type.
// This will be used to resolve this register in service container
type ControllerRegister []ControllerRegisterItem

// ControllerRegisterItem defines a specific controller
type ControllerRegisterItem struct {
	Controller interface{}
}
