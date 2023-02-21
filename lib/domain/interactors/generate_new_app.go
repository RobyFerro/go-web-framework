package interactors

import (
	"log"

	"github.com/RobyFerro/go-web-framework/lib/domain/services"
)

// GenerateNewApp generates new GoWeb application
type GenerateNewApp struct {
	Service     services.GitServices
	Destination string
	Repository  string
}

// Call execute Generate new app interactor
func (c GenerateNewApp) Call() {
	if err := c.Service.Clone(c.Destination, c.Repository); err != nil {
		log.Fatalf("Error: %s", err)
	}

	if err := c.Service.Remove(c.Destination); err != nil {
		log.Fatalf("Error: %s", err)
	}

	if err := c.Service.Update(c.Repository, c.Destination); err != nil {
		log.Fatalf("Error: %s", err)
	}
}
