package registers

// MiddlewareRegister defines all middleware present in your web application
type MiddlewareRegister []interface{}

// Add new middleware to Middleware register
func (c *MiddlewareRegister) Add(middleware interface{}) {
	*c = append(*c, middleware)
}
