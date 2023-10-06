package main

import (
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
)

const url = "http://crls.pki.goog/gts1c3/zdATt0Ex_Fk.crl"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error fetching:", err.Error())
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading:", err.Error())
	}
	resp.Body.Close()

	fmt.Println(x509.ParseRevocationList(body))
}
