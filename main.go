package main

import (
	"fmt"
	"gPhotosToFlickr/config"
	"gPhotosToFlickr/routehandler/googlehandler"
	"gPhotosToFlickr/routehandler/webhandler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// AppConfig is the app's configuration
var AppConfig config.Config

func main() {
	config.ReadConfig(&AppConfig)

	router := mux.NewRouter().StrictSlash(true)
	initHandlers(router)

	fmt.Println("All routes registered!")
	fmt.Println("Listening on port " + AppConfig.Server.Port)

	log.Fatal(http.ListenAndServe(":"+AppConfig.Server.Port, router))
}

// initHandlers initializes API and web handlers/routing
func initHandlers(router *mux.Router) {
	googlehandler.AppConfig = &AppConfig
	googlehandler.RegisterRoutes(router)

	webhandler.RegisterRoutes(router)
}
