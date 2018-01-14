package excel

import (
	"io"
	"net/http"
	"os"

	"github.com/sudhanshuraheja/ifsc/config"
	"github.com/sudhanshuraheja/ifsc/logger"
)

// DownloadLatest fetches the latest file from the server
func DownloadLatest() error {
	config.Init()
	logger.Init()

	downloadPath := "./data/latestExcel.xlsx"
	err := download(config.LatestDataExcel(), downloadPath)
	if err != nil {
		logger.Fatalln("Could not downlaod file", config.LatestDataExcel(), "to", downloadPath, "because of", err.Error())
	}
	return nil
}

// DownloadFile allows you to download the latest excel file
func download(url string, path string) error {
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
	logger.Infoln("Successfully downloaded file", url, "to", path)
	return nil
}
