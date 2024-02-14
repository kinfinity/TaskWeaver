package node

import (
	"net/http"
	"taskweaver/api/middleware"
	"taskweaver/api/rest"

	"github.com/gorilla/mux"
)

type NodeRouter struct {
	baseRouter     *mux.Router
	subRouter      *mux.Router
	routes         []rest.Route
	authMiddleware middleware.AuthMiddleware
}

func (n *NodeRouter) RegisterRoutes(route string, handler *http.HandlerFunc) *NodeRouter {
	return n
}

func (n NodeRouter) AddMiddleware() []func(http.Handler) http.Handler { return nil }

// NewNodeRouter returns a new instance of the router
func NewNodeRouter(baseRouter *mux.Router, RouterPath string, authMw middleware.AuthMiddleware) *NodeRouter {

	nodeRouter := &NodeRouter{
		baseRouter:     baseRouter,
		subRouter:      baseRouter.PathPrefix(RouterPath).Subrouter(),
		authMiddleware: authMw,
	}
	nodeRouter.setupRoutes()
	return nodeRouter
}

func (n *NodeRouter) setupRoutes() {
	// Create a map of paths to handler functions
	n.routes = []rest.Route{
		rest.NewRoute("GET", "/api/v1/health", HealthCheck),
		rest.NewRouteWithMiddleware("POST", "/api/v1/nodes", AddNode, n.authMiddleware.HandleMiddleware()),
		rest.NewRouteWithMiddleware("DELETE", "/api/v1/nodes/{id}", RemoveNode, n.authMiddleware.HandleMiddleware()),
		rest.NewRoute("GET", "/api/v1/nodes", GetNodes),
		rest.NewRoute("GET", "/api/v1/nodes/{id}", GetNode),
	}
}
