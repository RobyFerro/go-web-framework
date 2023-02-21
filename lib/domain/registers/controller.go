package registers

// ControllerRegister defines a controller register type.
// This will be used to resolve this register in service container
type ControllerRegister []interface{}

// Add a new controller to Controller register
func (c *ControllerRegister) Add(contoller interface{}) {
	*c = append(*c, contoller)
}
