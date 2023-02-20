package services

import (
	"net/http"

	"github.com/RobyFerro/go-web-framework/domain/entities"
)

// RouterService handles http router method
type RouterService interface {
	NewRouter(register entities.RouterRegister) http.Handler
}
