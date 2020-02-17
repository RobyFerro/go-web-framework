package gwf

import (
	"net/http"
)

// Main controller structure. This structure determines what you can find in the controllers instance.
// Adding something else inside this struct will not directly implement the struct. This because is just a part of the
// controller construction. See the "setBaseController" method inside app/kernel/kernel.go
type BaseController struct {
	Response http.ResponseWriter // HTTP response
	Request  *http.Request       // HTTP request
}
