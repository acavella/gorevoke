package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/viper"
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

	caid := viper.GetStringSlice("ca.id")
	cauri := viper.GetStringSlice("ca.uri")

	for i := 0; i < len(caid); i++ {

		err := DownloadFile(tmploc+caid[i]+".crl", cauri[i])
		if err != nil {
			fmt.Println("Error downloading file: ", err)
			return
		}
		fmt.Println("Downloaded: " + cauri[i])
		fmt.Println("Saved: " + tmploc + caid[i] + ".crl")
	}

	fmt.Println("Array length: ", len(caid))

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./crl/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//mux.HandleFunc("/", home)
	//mux.HandleFunc("/snippet/view", snippetView)
	//mux.HandleFunc("/snippet/create", snippetCreate)

	errhttp := http.ListenAndServe(":4000", mux)
	fmt.Println("Http error: ", errhttp)

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
