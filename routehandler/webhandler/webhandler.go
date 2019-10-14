package webhandler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/dist")))
}
