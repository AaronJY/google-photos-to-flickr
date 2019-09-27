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
	config.ReadEnv(&AppConfig)

	router := mux.NewRouter().StrictSlash(true)

	googleHandler.AppConfig = &AppConfig
	googleHandler.RegisterRoutes(router)

	fmt.Println("All routes registered!")
	fmt.Println("Listening!")

	log.Fatal(http.ListenAndServe(":1337", router))
}
