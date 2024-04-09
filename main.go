package main

import (
	"crypto/x509"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var workpath = "/usr/local/bin/gorevoke"

var appVersion = "0.0.0"
var appBuild = "UNK"
var appBuildDate = "00000000-0000"

func init() {
	log.SetFormatter(&log.TextFormatter{
		QuoteEmptyFields: true,
		FullTimestamp:    true,
	})

	directory, err := filepath.Abs(filepath.Dir(os.Args[0])) //get the current working directory
	if err != nil {
		log.Fatal(err) //print the error if obtained
	}

	workpath = directory // set app working directory

	viper.SetConfigName("config")            // name of config file (without extension)
	viper.SetConfigType("yaml")              // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./conf/")           // optionally look for config in the working directory
	viper.AddConfigPath(workpath + "/conf/") // optionally look for config in the working directory

	err2 := viper.ReadInConfig() // Find and read the config file
	if err2 != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	printver()

}

func main() {

	// Sets CA ID and URI to string arrays
	caid := viper.GetStringSlice("ca.id")
	cauri := viper.GetStringSlice("ca.uri")
	refresh := viper.GetInt("default.interval")
	webport := viper.GetString("default.port")
	server := viper.GetBool("default.webserver")

	if server {
		go webserver(webport)
	}

	log.Info("CRLs in list: ", len(caid))
	log.Info("Refresh interval: ", time.Duration(int(time.Second)*int(refresh)))

	//getcrl(caid, cauri, refresh)

	for {
		for i := 0; i < len(caid); i++ {

			var tmpfile string = workpath + "/crl/tmp/" + caid[i] + ".crl"
			var httpfile string = workpath + "/crl/static/" + caid[i] + ".crl"

			DownloadFile(tmpfile, cauri[i]) // Download CRL from remote

			crlfile, err := os.ReadFile(tmpfile)
			if err != nil {
				log.Error("Problem opening downloaded file: ", err)
				log.Info("Moving to next CRL entry.")
				goto SKIP
			} else {
				crl, err := x509.ParseRevocationList(crlfile)
				if err != nil {
					log.Errorln("Skipping CRL: ", err)
					goto SKIP
				} else {
					log.Infof("CRL %s is valid, issued by %s", caid[i], crl.Issuer.CommonName)
				}
			}

			if _, err := os.Stat(httpfile); err == nil {
				// file exists
				log.Info("CRL already exists")
				h1, err := getHash(tmpfile)
				if err != nil {
					log.Error("Error hashing: ", err)
				}
				h2, err2 := getHash(httpfile)
				if err2 != nil {
					log.Error("Error hashing: ", err2)
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
				log.Info("CRL is new, copying to: ", httpfile)
				copy(tmpfile, httpfile)
			} else {
				// catch anything else
				return
			}
		SKIP:
		}
		time.Sleep(time.Duration(int(time.Second) * refresh)) // Defines time to sleep before repeating
	}

}

func printver() {
	fmt.Printf("GoRevoke ver. %s\n", appVersion)
	fmt.Printf("Build Type: %s\n", appBuild)
	fmt.Printf("Build Date: %s\n", appBuildDate)
}
