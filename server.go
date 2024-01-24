package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func webserver(webport string) {
	// Disabled for testing
	// Simple http fileserver, serves all files in ./crl/static/
	// via localhost:4000/static/filename
	log.Info("Webserver is starting on port ", webport)
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(workpath + "/crl/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	errhttp := http.ListenAndServe(":"+webport, mux)
	log.Error("Http error: ", errhttp)
}
