package middleware

import (
	"log"
	"net/http"
)

// Create similar logging across requests
type LoggerMiddleware struct {
	logger *log.Logger
}

// middleware function to log HTTP requests
func (lm *LoggerMiddleware) HandleMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Log request details
			lm.logger.Printf("%s %s %s\n", r.Method, r.URL.Path, r.RemoteAddr)

			// Call the next handler in the chain
			next.ServeHTTP(w, r)
		})
	}
}

// NewLogger creates a new instance of the logger middleware
func NewRequestLogger(logger *log.Logger) *LoggerMiddleware {
	// logger.SetPrefix("[REQUEST]: ")
	return &LoggerMiddleware{logger: logger}
}
