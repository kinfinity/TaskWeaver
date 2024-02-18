package middleware

import (
	"net/http"
)

// Create similar logging across requests
type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

// Authenticate is the middleware function for authentication
func (am *AuthMiddleware) HandleMiddleware() MiddlewareFunc {
	return func(next http.Handler) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract the authentication token from the request header
			token := r.Header.Get("Authorization")

			// Check if the token is valid (you should implement your own validation logic)
			if isValidToken(token) {
				// If valid, call the next handler in the chain
				next.ServeHTTP(w, r)
			} else {
				// If not valid, respond with unauthorized status
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}
		})
	}
}

// isValidToken is a placeholder function for token validation (implement your own logic)
func isValidToken(token string) bool {
	// Implement your token validation logic (e.g., check against a database or third-party service)
	// For simplicity, this placeholder always returns true
	return true
}
