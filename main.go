package main

import (
	"fmt"
	"gPhotosToFlickr/routeHandlers/googleHandler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	googleHandler.RegisterRoutes(router)

	fmt.Println("All routes registered!")
	fmt.Println("Listening!")

	log.Fatal(http.ListenAndServe(":1337", router))
}