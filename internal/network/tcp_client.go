package network

import (
	"fmt"
	"net"
	"time"
)

type BackendClient struct {
	Conn net.Conn
	Addr string
}

var Backend *BackendClient

func InitBackendConnect(address string) {
	Backend = &BackendClient{Addr: address}
	go Backend.reconnectLoop()
}

func (b *BackendClient) reconnectLoop() {
	for {
		if b.Conn == nil {
			conn, err := net.DialTimeout("tcp", b.Addr, 5*time.Second)
			if err != nil {
				fmt.Printf("Échec connexion C++ (%s), nouvel essai dans 3s...\n", err)
				time.Sleep(3 * time.Second)
				continue
			}
			fmt.Println("Connecté au serveur C++ avec succès ✅")
			b.Conn = conn
		}
		time.Sleep(5 * time.Second)
	}
}

func (b *BackendClient) Send(data []byte) error {
	if b.Conn == nil {
		return fmt.Errorf("backend non connecté")
	}
	_, err := b.Conn.Write(data)
	return err
}
