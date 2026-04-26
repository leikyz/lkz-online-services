package handlers

import (
	"fmt"
	"net"

	"github.com/leikyz/lkz-online-services/internal/models"
	"github.com/leikyz/lkz-online-services/internal/network"
	"github.com/leikyz/lkz-online-services/internal/network/messages/approach"
	"github.com/leikyz/lkz-online-services/internal/network/messages/lobbies"
	"github.com/leikyz/lkz-online-services/internal/network/messages/metrics"
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

func HandlePingPong(msg *metrics.PingPongMessage, c *models.Client, conn net.Conn) (*models.Client, error) {
	// Create a new response message
	response := metrics.NewPingPongMessage()

	// CRITICAL: Echo the sequence ID back to the client
	response.SequenceID = msg.SequenceID

	data, err := response.Serialize()
	if err != nil {
		return c, fmt.Errorf("ping serialization failed: %v", err)
	}

	// Use the client's connection to send it back
	_, err = c.Conn.Write(data)
	if err != nil {
		return c, fmt.Errorf("failed to send ping response: %v", err)
	}

	return c, nil
}

func HandleClientInGameHandShake(msg *metrics.ClientInGameHandShakeMessage, c *models.Client, conn net.Conn) (*models.Client, error) {
	fmt.Printf("Received in-game handshake from client %d\n", c.ID)

	s, _ := registries.Sessions.GetByLobby(c.Lobby)

	message := metrics.NewBackendMetricsMessage(c.Lobby.ID, uint16(c.Lobby.GetClientCount()), s.ID)
	data, _ := message.Serialize()

	c.Conn.Write(data)

	return c, nil
}

func HandleChangeReadyStatus(msg *lobbies.ChangeReadyStatusMessage, c *models.Client, conn net.Conn) (*models.Client, error) {
	if c.Lobby == nil {
		return c, nil
	}

	c.Ready = !c.Ready
	var isEveryonReady bool = true

	for senderIndex, client := range c.Lobby.Clients {
		if client.ID == c.ID {
			client.Ready = c.Ready

			message := lobbies.NewChangeReadyStatusMessage(client.Ready, uint8(senderIndex))
			data, _ := message.Serialize()

			for _, clientToSend := range c.Lobby.Clients {

				if !clientToSend.Ready {
					isEveryonReady = false
				}

				clientToSend.Conn.Write(data)
			}
		}
	}

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
