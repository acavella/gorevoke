package main

import (
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

const tmploc = "./crl/tmp/"

var appVersion = "v0.0.0"
var appBuild = "0000000"

func init() {

	viper.SetConfigName("config")        // name of config file (without extension)
	viper.SetConfigType("yaml")          // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appname/") // path to look for the config file in
	viper.AddConfigPath("./conf/")       // optionally look for config in the working directory

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	log.SetFormatter(&log.TextFormatter{
		QuoteEmptyFields: true,
		FullTimestamp:    true,
	})
	printver()

}

func main() {

	// Sets CA ID and URI to string arrays
	caid := viper.GetStringSlice("ca.id")
	cauri := viper.GetStringSlice("ca.uri")
	refresh := viper.GetInt("default.interval")
	webport := viper.GetString("default.port")

	go webserver(webport)

	log.Info("CRLs in list: ", len(caid))
	log.Info("Refresh interval: ", time.Duration(int(time.Second)*int(refresh)))

	getcrl(caid, cauri, refresh)

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

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func getcrl(caid []string, cauri []string, refresh int) {
	for {
		log.Info("Checking for new CRL(s)")
		// Simple loop through arrays, downloads each crl from source
		for i := 0; i < len(caid); i++ {

			var tmpfile string = tmploc + caid[i] + ".crl"
			var httpfile string = "./crl/static/" + caid[i] + ".crl"

			err := DownloadFile(tmpfile, cauri[i])
			if err != nil {
				fmt.Println("Error downloading file: ", err)
				return
			}
			log.Info("Downloading file: ", cauri[i])
			log.Info("Download location: ", tmpfile)

			if _, err := os.Stat(httpfile); err == nil {
				// file exists
				h1, err := getHash(tmpfile)
				if err != nil {
					log.Error("Error hashing: ", err)
					return
				}
				h2, err2 := getHash(httpfile)
				if err2 != nil {
					log.Error("Error hashing: ", err2)
					return
				}
				log.Debug(h1, h2, h1 == h2)
				if h1 != h2 {
					log.Info("File hashes do not match: ", h1, h2)
					log.Info("Copying file to destination: ", httpfile)
					copy(tmpfile, httpfile)
				} else {
					log.Info("No changes detected, proceeding.")
				}
			} else if errors.Is(err, os.ErrNotExist) {
				// file does not exist
				log.Info("Copying file to destination: ", httpfile)
				copy(tmpfile, httpfile)
			} else {
				// catch anything else
				return
			}

		}
		time.Sleep(time.Duration(int(time.Second) * refresh)) // Defines time to sleep before repeating
	}
}

func webserver(webport string) {
	// Disabled for testing
	// Simple http fileserver, serves all files in ./crl/static/
	// via localhost:4000/static/filename
	log.Info("Webserver is started on port ", webport)
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./crl/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	errhttp := http.ListenAndServe(":"+webport, mux)
	log.Error("Http error: ", errhttp)
}

func printver() {
	fmt.Printf("GoRevoke %s\n", appVersion)
	fmt.Printf("Build Number: %s\n", appBuild)
}
