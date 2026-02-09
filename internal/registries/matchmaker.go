package registries

import (
	"fmt"
	"sync"
	"time"

	"github.com/leikyz/lkz-online-services/internal/models"
	"github.com/leikyz/lkz-online-services/internal/network/messages/lobbies"
)

type Matchmaker struct {
	queue []*models.Client
	mu    sync.RWMutex
}

var Matchmaking = &Matchmaker{queue: make([]*models.Client, 0)}

func (m *Matchmaker) AddClientToQueue(client *models.Client) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.queue = append(m.queue, client)
	fmt.Printf("Client %s added to matchmaking queue\n", client.ID)
}

func (m *Matchmaker) FindMatches() {
	m.mu.Lock()
	// Check if there are any clients in the queue
	if len(m.queue) == 0 {
		m.mu.Unlock()
		return
	}

	// Pop the first client from the queue (FIFO)
	clientInComming := m.queue[0]
	m.queue = m.queue[1:] // Remove it from the queue
	m.mu.Unlock()

	var lobby *models.Lobby
	// Find a lobby suitable for this client
	Lobbies.mu.RLock()
	for _, lobbyCreated := range Lobbies.lobbies {
		// Ensure we pass 'client' to IsAvailable, not 'lobby'
		if lobbyCreated.IsAvailable(clientInComming) && !lobbyCreated.IsFull() {
			lobby = lobbyCreated
			break
		}
	}
	Lobbies.mu.RUnlock()

	if lobby != nil {
		lobby.Mu.Lock()
		// Add the 'client' to the lobby (not the lobby to itself)
		lobby.Clients = append(lobby.Clients, clientInComming)
		lobby.Mu.Unlock()
		fmt.Printf("Client %s added to lobby %s\n", clientInComming.ID, lobby.ID)
	} else {
		lobby = Lobbies.CreateLobby(0, 4)

		if lobby == nil {
			fmt.Println("[ERREUR] CreateLobby a renvoyé nil")
			return
		}

		lobby.Mu.Lock()
		lobby.Clients = append(lobby.Clients, clientInComming)
		lobby.Mu.Unlock()
		fmt.Printf("Client %s added to new lobby %s\n", clientInComming.ID, lobby.ID)
	}

	clientInComming.Lobby = lobby

	if lobby != nil {
		// We get the new client index
		var newClientIndex int
		for i, c := range lobby.Clients {
			if c == clientInComming {
				newClientIndex = i
				break
			}
		}

		for i, client := range lobby.Clients {
			if client == clientInComming {
				msg := lobbies.NewJoinLobbyMessage(uint8(i))
				data, _ := msg.Serialize()
				client.Conn.Write(data)
			} else {
				msg := lobbies.NewJoinLobbyMessage(uint8(newClientIndex))
				data, _ := msg.Serialize()
				client.Conn.Write(data)

				msgOld := lobbies.NewJoinLobbyMessage(uint8(i))
				dataOld, _ := msgOld.Serialize()
				clientInComming.Conn.Write(dataOld)
			}
		}
	}
}

func (m *Matchmaker) Start() {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		defer ticker.Stop()
		fmt.Println("Matchmaker loop started...")

		for range ticker.C {
			m.FindMatches()
		}
	}()
}
