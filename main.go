package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"net/http"
	"os"

	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

const tmploc = "./crl/tmp/"

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

	// Sets CA ID and URI to string arrays
	caid := viper.GetStringSlice("ca.id")
	cauri := viper.GetStringSlice("ca.uri")

	// Simple loop through arrays, downloads each crl from source
	for i := 0; i < len(caid); i++ {

		err := DownloadFile(tmploc+caid[i]+".crl", cauri[i])
		if err != nil {
			fmt.Println("Error downloading file: ", err)
			return
		}
		log.Info("Downloading file: ", cauri[i])
		log.Info("Download location: ", tmploc+caid[i]+".crl")
	}

	fmt.Println("Array length: ", len(caid))

	// Simple http fileserver, serves all files in ./crl/static/
	// via localhost:4000/static/filename
	/* Disabled for testing
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./crl/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	errhttp := http.ListenAndServe(":4000", mux)
	fmt.Println("Http error: ", errhttp)
	*/

	// Simple hash comparison
	h1, err := getHash("./crl/tmp/x21.crl")
	if err != nil {
		return
	}
	h2, err2 := getHash("./crl/static/x21.crl")
	if err2 != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)

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

func getHash(filename string) (uint32, error) {
	// open the file
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	// remember to always close opened files
	defer f.Close()

	// create a hasher

	h := crc32.NewIEEE()
	// copy the file into the hasher
	// - copy takes (dst, src) and returns (bytesWritten, error)
	_, err = io.Copy(h, f)
	// we don't care about how many bytes were written, but we do want to
	// handle the error
	if err != nil {
		return 0, err
	}
	return h.Sum32(), nil
}
