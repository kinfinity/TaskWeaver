package middleware

import "net/http"

type Middleware interface {
	HandleMiddleware() MiddlewareFunc
}

// Signature for Middleware
type MiddlewareFunc func(http.Handler) http.HandlerFunc
