package config

import (
	"github.com/AaronJY/google-photos-to-flickr/common/google"
	"github.com/AaronJY/google-photos-to-flickr/routehandler/googledownloadhandler/downloadinfo"
)

type AppState struct {
	DownloadInfo downloadinfo.DownloadInfo
	IsDownloading bool
	GoogleAuthToken google.AuthToken
}
