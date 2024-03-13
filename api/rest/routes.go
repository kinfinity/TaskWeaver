package rest

import (
	"net/http"
	"taskweaver/api/middleware"

	"github.com/gorilla/mux"
)

type Router interface {
	setupRoutes()
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

func Routes2MuxRouter(mr *mux.Router, routes []Route) {
	// Add routes to the Router
	for _, route := range routes {
		// create new route on mux router
		muxRoute := mr.NewRoute().Methods(string(route.Verb)).Path(route.Path)
		// If there are middleware functions defined for this route add them
		if len(route.middleware) > 0 {
			handler := route.Handler
			for i := len(route.middleware) - 1; i >= 0; i-- {
				handler = route.middleware[i](handler)
			}
			muxRoute.Handler(handler)
		} else {
			muxRoute.HandlerFunc(route.Handler)
		}
	}

}
