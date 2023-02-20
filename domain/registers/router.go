package registers

import "github.com/RobyFerro/go-web-framework/domain/entities"

// RouterRegister defines a new GoWeb Router. It contains all application routes.
type RouterRegister []entities.Router

// Add to router register
func (c RouterRegister) Add(router entities.Router) {
	c = append(c, router)
}
