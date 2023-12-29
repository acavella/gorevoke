package main

import (
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
)



func main() {

	fileUrl := "http://crls.pki.goog/gts1c3/zdATt0Ex_Fk.crl"
	
	err := DownloadFile("saveas.crl", fileUrl)
	if err != nil {
		fmt.Println("Error downloading file: ", err)
		return
}

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