package middlewares

import (
	"fmt"
	"net/http"
)

type LoggingMiddleware struct {
	Name        string
	Description string
}

// Handle set a limit of request allowed in a specific time
func (LoggingMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)

		next.ServeHTTP(w, r)
	})
}

// GetName returns the middleware name
func (m LoggingMiddleware) GetName() string {
	return m.Name
}

// GetDescription returns the middleware description
func (m LoggingMiddleware) GetDescription() string {
	return m.Description
}

func NewLoggingMiddleware() LoggingMiddleware {
	return LoggingMiddleware{
		Name:        "LoggingMiddleware",
		Description: "Prints to stdout incoming http requests",
	}
}
