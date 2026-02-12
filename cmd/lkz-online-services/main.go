package main

import (
	"fmt"

	"github.com/leikyz/lkz-online-services/internal/handlers"
	"github.com/leikyz/lkz-online-services/internal/network"
	"github.com/leikyz/lkz-online-services/internal/network/messages/approach"
	"github.com/leikyz/lkz-online-services/internal/network/messages/lobbies"
	"github.com/leikyz/lkz-online-services/internal/registries"
)

func main() {

	network.Register(1, func() network.Message { return approach.NewCreateClientMessage() }, network.Bind(handlers.HandleCreateClient))
	network.Register(4, func() network.Message { return &approach.StartMatchmakingMessage{} }, network.Bind(handlers.HandleStartMatchmaking))
	network.Register(6, func() network.Message { return &lobbies.ChangeReadyStatusMessage{} }, network.Bind(handlers.HandleChangeReadyStatus))

	if registries.Matchmaking == nil {
		fmt.Println("Erreur: Registre Matchmaking non initialisé")
	}
	go registries.Matchmaking.Start()

	fmt.Println("Services démarrés...")

	fmt.Println("Serveur en attente sur 127.0.0.1:8081...")
	network.Initialization("127.0.0.1:8081")

	select {}
}
