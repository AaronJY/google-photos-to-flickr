package routeHandlers

import "github.com/gorilla/mux"

type RouteHandler interface {
	RegisterRoutes(router *mux.Router)
}
