package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	fileUrl := "http://crls.pki.goog/gts1c3/zdATt0Ex_Fk.crl"
	savloc := "./crl/x21.crl"

	// Download the file, params:
	// 1) name of file to save as
	// 2) URL to download FROM
	err := DownloadFile(savloc, fileUrl)
	if err != nil {
		fmt.Println("Error downloading file: ", err)
		return
	}

	fmt.Println("Downloaded: " + fileUrl)
}

// DownloadFile will download from a given url to a file. It will
// write as it downloads (useful for large files).
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
