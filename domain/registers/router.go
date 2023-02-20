package registers

import "github.com/RobyFerro/go-web-framework/domain/entities"

// RouterRegister defines a new GoWeb Router. It contains all application routes.
type RouterRegister struct {
	Route  []entities.Route
	Groups []entities.Group
}
