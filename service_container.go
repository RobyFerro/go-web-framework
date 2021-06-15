package gwf

import (
	"github.com/RobyFerro/go-web-framework/foundation"
	"github.com/RobyFerro/go-web-framework/helper"
	"github.com/RobyFerro/go-web-framework/types"
	"go.uber.org/dig"
)

// BuildContainer provide the global service container
func BuildContainer() *dig.Container {
	container := dig.New()

	for _, s := range foundation.Services.List {
		if err := container.Provide(s); err != nil {
			helper.ProcessError(err)
		}
	}

	injectBasicEntities(container)
	foundation.Container = container

	return container
}

// Inject base entities: controllers, models, commands in service container
func injectBasicEntities(sc *dig.Container) {
	_ = sc.Provide(func() types.ControllerRegister {
		return foundation.Controllers
	})

	_ = sc.Provide(func() types.CommandRegister {
		return foundation.Commands
	})

	_ = sc.Provide(func() types.ModelRegister {
		return foundation.Models
	})
}
