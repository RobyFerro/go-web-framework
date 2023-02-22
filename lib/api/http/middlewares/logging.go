package middlewares

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware prints http requests to stdout
type LoggingMiddleware struct {
	Name        string
	Description string
}

// Handle set a limit of request allowed in a specific time
func (LoggingMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)

		log.Printf("%s %s -> Time: %v", r.Method, r.URL.Path, time.Since(start))
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

// NewLoggingMiddleware creates an instance LoggingMiddleware
func NewLoggingMiddleware() LoggingMiddleware {
	return LoggingMiddleware{
		Name:        "LoggingMiddleware",
		Description: "Prints to stdout incoming http requests",
	}
}
