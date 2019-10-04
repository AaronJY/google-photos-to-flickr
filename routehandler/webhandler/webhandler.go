package webhandler

import (
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes sets up routing for the web application
func RegisterRoutes(router *mux.Router) {
	router.Handle("/", http.FileServer(http.Dir("./public")))
}
