package main

import (
	"fmt"
	"gPhotosToFlickr/routeHandlers/googleHandler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	googleHandler.RegisterRoutes(router)

	fmt.Println("All routes registered!")
	fmt.Println("Listening!")

	log.Fatal(http.ListenAndServe(":1337", router))
}