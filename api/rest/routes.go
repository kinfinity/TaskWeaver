package rest

import (
	"net/http"
	"taskweaver/api/middleware"
)

type Router interface {
	SetupRoutes()
}

type Route struct {
	Verb       HTTPVerb
	Path       string
	Handler    http.HandlerFunc
	middleware []middleware.MiddlewareFunc
}

func NewRoute(method, path string, handler http.HandlerFunc) Route {
	Verb, _ := NewHTTPVerb(method)
	return Route{
		Verb:    Verb,
		Path:    path,
		Handler: handler,
	}
}

func NewRouteWithMiddleware(method, path string, handler http.HandlerFunc, middlewares ...middleware.MiddlewareFunc) Route {
	return NewRoute(method, path, handler).WithMiddlewares(middlewares...)
}

func (r Route) WithMiddlewares(mw ...middleware.MiddlewareFunc) Route {
	r.middleware = append(r.middleware, mw...)
	return r
}
