package excel

import (
	"io"
	"net/http"
	"os"
)

// DownloadFile allows you to download the latest excel file
func DownloadFile(url string, path string) error {
	// Create an empty file
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	// Download the file
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
