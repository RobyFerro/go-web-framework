package services

import (
	"net/http"

	register "github.com/RobyFerro/go-web-framework/domain/registers"
)

// RouterService handles http router method
type RouterService interface {
	NewRouter(register []register.RouterRegister) http.Handler
}
