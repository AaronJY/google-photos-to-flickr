package main

import (
	"fmt"
	"gPhotosToFlickr/config"
	"gPhotosToFlickr/routeHandlers/googleHandler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var AppConfig config.Config

func main() {
	config.ReadConfig(&AppConfig)

	router := mux.NewRouter().StrictSlash(true)

	googleHandler.AppConfig = &AppConfig
	googleHandler.RegisterRoutes(router)

	fmt.Println("All routes registered!")
	fmt.Println("Listening on port " + AppConfig.Server.Port)

	log.Fatal(http.ListenAndServe(":" + AppConfig.Server.Port, router))
}
