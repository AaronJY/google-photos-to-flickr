package googledownloadhandler

import (
	. "github.com/AaronJY/google-photos-to-flickr/config"
	"github.com/gorilla/mux"
	"net/http"
)

var AppConfig Config

var IsDownloading bool = false

const routePrefix = "/api/google/download"

// RegisterRoutes registers API routes for googledownloadhandler
func RegisterRoutes(router *mux.Router) {
	//subRouter := router.PathPrefix(routePrefix).Subrouter()
	//
	//fmt.Println("Successfully registered Google downloader routes.")
}

// Download initiates the download process
func Download() {

}

// DownloadStatus returns the status of the current download task
func GetDownloadStatus() {

}

// DownloadInfo returns detailed information on a completed download
func GetDownloadInfo(resp http.ResponseWriter, req *http.Request) {

}