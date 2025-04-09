package main

import (
	"crypto/x509"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

	printver()

	directory, err := filepath.Abs(filepath.Dir(os.Args[0])) //get the current working directory
	if err != nil {
		log.Fatal(err) //print the error if obtained
	}

	workpath = directory // set app working directory

	viper.SetConfigName("gorevoke")            // name of config file (without extension)
	viper.SetConfigType("yaml")              // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(workpath + "/conf/") // optionally look for config in the working directory
	viper.AddConfigPath("$HOME/.gorevoke/")
	viper.AddConfigPath("/etc/")

	// Set defaults
	viper.SetDefault("default.interval", 5)
	viper.SetDefault("default.webserver", false)
	viper.SetDefault("default.port", 4000)

	// Enable environment variable configuration
	viper.AutomaticEnv()
	viper.SetEnvPrefix("gorevoke")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if configErr := viper.ReadInConfig(); configErr != nil {
		if _, ok := configErr.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
			log.Warn("no config file found, using defaults/environment")
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("fatal error reading config file: %w", configErr))
		}
	}
}

func main() {
	// Retrieve config values
	crls := viper.GetStringMapString("crls")
	refresh := viper.GetInt("default.interval")
	webport := viper.GetString("default.port")
	server := viper.GetBool("default.webserver")

	if server {
		go webserver(webport)
	}

	log.Info("CRLs in list: ", len(crls))
	log.Info("Refresh interval: ", time.Duration(int(time.Second)*int(refresh)))

	for {
		for caId, caUrl := range crls {
			var tmpfile string = workpath + "/crl/tmp/" + caId + ".crl"
			var httpfile string = workpath + "/crl/static/" + caId + ".crl"

			DownloadFile(tmpfile, caUrl) // Download CRL from remote

			crlfile, err := os.ReadFile(tmpfile)
			if err != nil {
				log.Error("Problem opening downloaded file: ", err)
				log.Info("Moving to next CRL entry.")
				continue
			} else {
				crl, err := x509.ParseRevocationList(crlfile)
				if err != nil {
					log.Errorln("Skipping CRL: ", err)
					continue
				} else {
					log.Infof("CRL %s is valid, issued by %s", caId, crl.Issuer.CommonName)
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
		}
		time.Sleep(time.Duration(int(time.Second) * refresh)) // Defines time to sleep before repeating
	}

}

func printver() {
	fmt.Printf("GoRevoke ver. %s\n", appVersion)
	fmt.Printf("Build Type: %s\n", appBuild)
	fmt.Printf("Build Date: %s\n", appBuildDate)
}
