package network

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"time"
)

type TCPClient struct {
	Conn net.Conn
	Addr string
}

type TCPServer struct {
	Listener net.Listener
}

var Client *TCPClient
var Server *TCPServer

func Initialization(address string) {

	// Client initialization to connect to C++ Server
	Client = &TCPClient{Addr: address}
	go Client.TryConnect()

	// TCP/TLS Server initialization to accept connections from game clients
	cert, err := tls.LoadX509KeyPair(".local/certs/cert.crt", ".local/certs/cert.key")
	if err != nil {
		log.Fatal(err)
	}

	// TLS configuration
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	ln, err := tls.Listen("tcp", ":8080", config)
	if err != nil {
		log.Fatal(err)
	}

	Server = &TCPServer{Listener: ln}

	go Server.Poll()

}

// Poll method to accept incoming connections and packet handling
func (b *TCPServer) Poll() {
	for {
		conn, err := b.Listener.Accept()
		if err != nil {
			continue
		}
		go HandleMessage(conn)
	}
}

// TryConnect method to establish and maintain connection to C++ Server with retries
func (b *TCPClient) TryConnect() {
	for {
		if b.Conn == nil {
			conn, err := net.DialTimeout("tcp", b.Addr, 5*time.Second)
			if err != nil {
				fmt.Printf("Connection failed (%s), next try in 3s...\n", err)
				time.Sleep(3 * time.Second)
				continue
			}
			fmt.Println("Connected to C++ Server successfuly")
			b.Conn = conn
		}
		time.Sleep(5 * time.Second)
	}
}

// Send method to transmit data to client
func (b *TCPClient) Send(data []byte) error {
	if b.Conn == nil {
		return fmt.Errorf("backend non connecté")
	}
	_, err := b.Conn.Write(data)
	return err
}
