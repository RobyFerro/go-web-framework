package services

import (
	"net/http"

	"github.com/RobyFerro/go-web-framework/lib/domain/registers"
)

// RouterService handles http router method
type RouterService interface {
	NewRouter(register registers.RouterRegister) http.Handler
}
