package handlers

import (
	"fmt"
	"net"

	"github.com/leikyz/lkz-online-services/internal/models"
	"github.com/leikyz/lkz-online-services/internal/network"
	"github.com/leikyz/lkz-online-services/internal/network/messages/approach"
	"github.com/leikyz/lkz-online-services/internal/network/messages/lobbies"
	"github.com/leikyz/lkz-online-services/internal/network/messages/sessions"
	"github.com/leikyz/lkz-online-services/internal/registries"
)

func HandleStartMatchmaking(msg *approach.StartMatchmakingMessage, c *models.Client, conn net.Conn) (*models.Client, error) {
	registries.Matchmaking.AddClientToQueue(c)

	return c, nil
}

func HandleCreateClient(msg *approach.CreateClientMessage, c *models.Client, conn net.Conn) (*models.Client, error) {
	fmt.Printf("Handling CreateClient for connection: %v\n", conn.RemoteAddr())

	clientPtr := registries.Clients.CreateClient("guest", 1, conn)

	message := approach.NewWelcomeMessage(clientPtr.ID)
	data, _ := message.Serialize()
	clientPtr.Conn.Write(data)

	return clientPtr, nil
}

func HandleChangeReadyStatus(msg *lobbies.ChangeReadyStatusMessage, c *models.Client, conn net.Conn) (*models.Client, error) {
	if c.Lobby == nil {
		return c, nil
	}

	c.Ready = !c.Ready
	var isEveryonReady bool = true

	for i, client := range c.Lobby.Clients {
		if !client.Ready {
			isEveryonReady = false
		}
		message := lobbies.NewChangeReadyStatusMessage(client.Ready, uint8(i))
		data, _ := message.Serialize()
		client.Conn.Write(data)
	}

	// Send to
	if isEveryonReady {
		session := registries.Sessions.CreateSession(c.Lobby)
		var allowedIDs []uint32

		for _, client := range c.Lobby.Clients {
			allowedIDs = append(allowedIDs, uint32(client.ID))
		}

		startGameMsg := sessions.NewCreateSessionMessage(session.ID, allowedIDs)
		data, _ := startGameMsg.Serialize()

		network.Client.Send(data)

		for _, client := range c.Lobby.Clients {
			joinSessionMsg := sessions.NewSessionAssignmentMessage(session.ID)
			data, _ := joinSessionMsg.Serialize()
			client.Conn.Write(data)
		}
	}

	return c, nil
}
