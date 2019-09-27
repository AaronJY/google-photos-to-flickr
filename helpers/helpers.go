package helpers

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

// openbrowser executes an OS command to open a browser
// at a given URL. Works cross-platform (linux, windows, mac)
func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}