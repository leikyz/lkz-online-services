package main

import (
	"fmt"

	"github.com/leikyz/lkz-online-services/internal/network"
	"github.com/leikyz/lkz-online-services/internal/network/messages/approach"
	"github.com/leikyz/lkz-online-services/internal/network/messages/lobbies"
	"github.com/leikyz/lkz-online-services/internal/registries"
)

func main() {
	network.RegisterMessage(1, func() network.Message { return approach.NewCreateClientMessage() })
	network.RegisterMessage(4, func() network.Message { return approach.NewStartMatchmakingMessage() })
	network.RegisterMessage(6, func() network.Message { return lobbies.NewChangeReadyStatusMessage(false, 0) })
	if registries.Matchmaking == nil {
		fmt.Println("Erreur: Registre Matchmaking non initialisé")
	}
	go registries.Matchmaking.Start()

	fmt.Println("Services démarrés...")

	fmt.Println("Serveur en attente sur 127.0.0.1:8081...")
	network.Initialization("127.0.0.1:8081")

	select {}
}
