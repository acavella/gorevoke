package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

func init() {

	viper.SetConfigName("config")        // name of config file (without extension)
	viper.SetConfigType("yaml")          // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appname/") // path to look for the config file in
	viper.AddConfigPath("./conf/")       // optionally look for config in the working directory

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

}

func main() {

	fileUrl := "http://crls.pki.goog/gts1c3/zdATt0Ex_Fk.crl"
	savloc := "./crl/x21.crl"

	ca := viper.GetStringMap("ca.1.id")
	//arraylen := len(ca)
	fmt.Println(ca)
	//fmt.Println(ca[0])
	//fmt.Println("Array length: ", arraylen)

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
