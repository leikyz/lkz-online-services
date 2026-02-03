package main

import (
	"crypto/tls"
	"log"

	"github.com/leikyz/lkz-online-services/internal/network"
)

func main() {
	network.InitBackendConnect("127.0.0.1:8081")

	cert, err := tls.LoadX509KeyPair(".local/certs/cert.crt", ".local/certs/cert.key")
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	ln, err := tls.Listen("tcp", ":8080", config)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Serveur Gateway TCP/TLS actif sur :8080 🚀")

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go network.HandleConnection(conn)
	}
}
