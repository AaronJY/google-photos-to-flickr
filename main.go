package main

import (
	"fmt"
	"github.com/AaronJY/google-photos-to-flickr/routehandler/googlehandler"
	"github.com/AaronJY/google-photos-to-flickr/routehandler/webhandler"
	"github.com/AaronJY/google-photos-to-flickr/config"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// appConfig is the app's configuration
var appConfig config.Config

func main() {
	// Load config
	config.ReadConfig(&appConfig)

	// Setup routers
	router := mux.NewRouter().StrictSlash(true)
	initHandlers(router)

	serverPortStr := strconv.Itoa(appConfig.Server.Port)
	fmt.Println("Listening on port " + serverPortStr)
	log.Fatal(http.ListenAndServe(":" + serverPortStr, router))
}

// initHandlers initializes API and web handlers/routing
func initHandlers(router *mux.Router) {
	googlehandler.AppConfig = &appConfig
	googlehandler.RegisterRoutes(router)

	webhandler.RegisterRoutes(router)

	fmt.Println("All routes registered")
}
