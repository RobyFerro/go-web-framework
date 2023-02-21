package registers

// ModelRegister defines a controller register type.
// This will be used to resolve this register in service container
type ModelRegister []interface{}

// Add new model to Model register
func (c *ModelRegister) Add(model interface{}) {
	*c = append(*c, model)
}
