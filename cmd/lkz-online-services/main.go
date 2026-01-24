package main

import (
	"fmt"
	"github.com/quic-go/quic-go/http3"
	"github.com/leikyz/lkz-online-services/internal/network"
)

func main() {
	network.RegisterHandlers()

	fmt.Println("LKZ Online Services Starting... 🚀")

	err := http3.ListenAndServeQUIC(":8080", ".local/certs/certs.crt", ".local/certs/cert.key", nil)

	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
