package middleware

import (
	"net/http"
)

type @@TMP@@Middleware struct {
    Name        string
    Description string
}

// Handle description
func (@@TMP@@Middleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
	})
}

// GetName returns the middleware name
func (m *@@TMP@@Middleware) GetName() string {
	return m.Name
}

// GetDescription returns the middleware description
func (m *@@TMP@@Middleware) GetDescription() string {
	return m.Description
}

func New@@TMP@@Middleware() @@TMP@@Middleware{
	return @@TMP@@Middleware{
		Name:        "REPLACE WITH YOUR MIDDLEWARE NAME",
		Description: "REPLACE WITH YOUR MIDDLEWARE DESCRIPTION",
	}
}

