package main

import (
	"fmt"

	"github.com/leikyz/lkz-online-services/internal/network"
	"github.com/leikyz/lkz-online-services/internal/network/messages/approach"
)

func main() {
	network.RegisterMessage(1, func() network.Message {
		return approach.NewCreateClientMessage()
	})

	network.Initialization("127.0.0.1:8081")
	fmt.Println("Serveur démarrée...")
	select {}
}
