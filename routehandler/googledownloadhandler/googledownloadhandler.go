package googledownloadhandler

import (
	"github.com/AaronJY/google-photos-to-flickr/config"
	"github.com/gorilla/mux"
	"net/http"
)

var AppConfig config.Config
var AppState config.AppState

const (
	routePrefix = "/api/google/download"
	googlePhotosApiBaseUrl = "https://photoslibrary.googleapis.com/v1/"
)

// RegisterRoutes registers API routes for googledownloadhandler
func RegisterRoutes(router *mux.Router) {
	//subRouter := router.PathPrefix(routePrefix).Subrouter()
	//
	//fmt.Println("Successfully registered Google downloader routes.")
}

// Download initiates the download process
func Download(resp http.ResponseWriter, req *http.Request) {
	AppState.IsDownloading = true


}

func doDownload(nextPageToken string) {
}

// DownloadInfo returns detailed information on a completed download
func GetDownloadInfo(resp http.ResponseWriter, req *http.Request) {

}