package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	. "github.com/AaronJY/google-photos-to-flickr/config"
	"github.com/AaronJY/google-photos-to-flickr/routehandler/googlehandler"
	"github.com/AaronJY/google-photos-to-flickr/routehandler/webhandler"
	"github.com/gorilla/mux"
)

var appConfig Config
var appState AppState

func main() {
	ReadConfig(&appConfig)

	router := mux.NewRouter().StrictSlash(true)
	initRouteHandlers(router)

	serverPortStr := strconv.Itoa(appConfig.Server.Port)
	fmt.Println("Listening on port " + serverPortStr)
	log.Fatal(http.ListenAndServe(":"+serverPortStr, router))
}

func initRouteHandlers(router *mux.Router) {
	googlehandler.AppConfig = &appConfig
	googlehandler.AppState = &appState
	googlehandler.RegisterRoutes(router)

	webhandler.RegisterRoutes(router)

	fmt.Println("All routes registered")
}
