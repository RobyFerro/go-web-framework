package foundation

import (
	"github.com/RobyFerro/go-web-framework/kernel"
	"go.uber.org/dig"
)

var singletonIOC *dig.Container

// RetrieveSingletonContainer returns a IOC container that contains every IOC singleton services.
func RetrieveSingletonContainer() *dig.Container {
	return singletonIOC
}

// RetrieveServiceContainer returns a IOC container that contains every standard IOC services
func RetrieveServiceContainer() *dig.Container {
	return kernel.BuildCustomContainer()
}
