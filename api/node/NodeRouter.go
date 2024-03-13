package node

import (
	"taskweaver/api/middleware"
	"taskweaver/api/rest"

	"github.com/gorilla/mux"
)

type NodeRouter struct {
	rest.Router
	Name           string
	baseRouter     *mux.Router
	subRouter      *mux.Router
	routes         []rest.Route
	authMiddleware middleware.AuthMiddleware
}

// NewNodeRouter returns a new instance of the router
func NewNodeRouter(baseRouter *mux.Router, RouterPath string, authMw middleware.AuthMiddleware) *NodeRouter {
	nodeRouter := &NodeRouter{
		baseRouter:     baseRouter,
		subRouter:      baseRouter.PathPrefix(RouterPath).Subrouter(),
		authMiddleware: authMw,
	}
	nodeRouter.setupRoutes()
	rest.Routes2MuxRouter(nodeRouter.subRouter, nodeRouter.routes)
	return nodeRouter
}

func (n *NodeRouter) setupRoutes() {
	// Create a map of paths to handler functions
	n.routes = []rest.Route{
		rest.NewRouteWithMiddleware("POST", "", AddNode, n.authMiddleware.HandleMiddleware()),
		rest.NewRouteWithMiddleware("DELETE", "/{id}", RemoveNode, n.authMiddleware.HandleMiddleware()),
		rest.NewRoute("GET", "", GetNodes),
		rest.NewRoute("GET", "/{id}", GetNode),
	}
}
