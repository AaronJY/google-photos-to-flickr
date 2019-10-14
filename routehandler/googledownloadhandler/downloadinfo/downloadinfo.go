package downloadinfo

type DownloadInfo struct {
	fileCount int
}

type DownloadStatus int

const (
	Unknown          = 0
	NotStarted       = 1
	Processing       = 2
	DownloadingMedia = 3
	Finished         = 4
)
