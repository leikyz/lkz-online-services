package main

import (
	"fmt"
	"net/http"

	"github.com/leikyz/lkz-online-services/internal/network"
)

func main() {
	mux := http.NewServeMux()
	network.RegisterHandlers(mux)

	fmt.Println("LKZ Online Services Starting on :8080... 🚀")
	fmt.Println("Listening for HTTP/2 (TCP)")

	// Serve HTTP/2 with TLS
	err := http.ListenAndServeTLS(":8080", ".local/certs/cert.crt", ".local/certs/cert.key", mux)

	if err != nil {
		fmt.Println("TCP Server Error:", err)
	}
}
