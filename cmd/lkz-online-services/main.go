package main

import (
	"fmt"

	"github.com/leikyz/lkz-online-services/internal/network"
	"github.com/leikyz/lkz-online-services/internal/network/messages/approach"
	"github.com/leikyz/lkz-online-services/internal/registries"
)

func main() {
	network.RegisterMessage(1, func() network.Message {
		return approach.NewCreateClientMessage()
	})

	network.RegisterMessage(4, func() network.Message {
		return approach.NewStartMatchmakingMessage()
	})

	network.Initialization("127.0.0.1:8081")
	fmt.Println("Serveur démarrée...")

	registries.Matchmaking.Start()
	select {}		
}
