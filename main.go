package main

import (
	"fmt"
	"gPhotosToFlickr/config"
	"gPhotosToFlickr/routehandler/googlehandler"
	"gPhotosToFlickr/routehandler/webhandler"
	"github.com/go-redis/redis"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// appConfig is the app's configuration
var appConfig config.Config
var redisClient redis.Client

func main() {
	// Load config
	config.ReadConfig(&appConfig)

	// Setup routers
	router := mux.NewRouter().StrictSlash(true)
	initHandlers(router)

	// Setup redis
	setupRedis()

	serverPortStr := strconv.Itoa(appConfig.Server.Port)
	fmt.Println("Listening on port " + serverPortStr)
	log.Fatal(http.ListenAndServe(":" + serverPortStr, router))
}

func setupRedis() {
	// Setup redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: appConfig.Redis.Address,
		Password: "",
		DB: 0, // use default database
	})

	pong, err := redisClient.Ping().Result()
	fmt.Println(pong, err)
}

// initHandlers initializes API and web handlers/routing
func initHandlers(router *mux.Router) {
	googlehandler.AppConfig = &appConfig
	googlehandler.RegisterRoutes(router)

	webhandler.RegisterRoutes(router)

	fmt.Println("All routes registered!")
}
