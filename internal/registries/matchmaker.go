package registries

import (
	"sync"
    "time"
    "fmt"

	"github.com/leikyz/lkz-online-services/internal/models"
)

type Matchmaker struct {
	queue []*models.Client
	mu      sync.RWMutex

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
    client := m.queue[0]
    m.queue = m.queue[1:] // Remove it from the queue
    m.mu.Unlock()

    var availableLobby *models.Lobby
    
    // Find a lobby suitable for this client
    Lobbies.mu.RLock()
    for _, lobby := range Lobbies.lobbies {
        // Ensure we pass 'client' to IsAvailable, not 'lobby'
        if lobby.IsAvailable(client) && !lobby.IsFull() {
            availableLobby = lobby
            break
        }
    }
    Lobbies.mu.RUnlock()

    if availableLobby != nil {
        availableLobby.Mu.Lock()
        // Add the 'client' to the lobby (not the lobby to itself)
        availableLobby.Clients = append(availableLobby.Clients, client)
        availableLobby.Mu.Unlock()
        fmt.Printf("Client %s added to lobby %s\n", client.ID, availableLobby.ID)
    } else {
        // Create a new lobby if none are available
        newLobby := Lobbies.CreateLobby(0, 4)
        newLobby.Mu.Lock()
        newLobby.Clients = append(newLobby.Clients, client)
        newLobby.Mu.Unlock()
        fmt.Printf("Client %s added to new lobby %s\n", client.ID, newLobby.ID)
    }
}

func (m *Matchmaker) Start() {
    // Set interval (e.g., every second)
    ticker := time.NewTicker(1 * time.Second)

    // Run the loop in a goroutine
    go func() {
        // Optional: ensure the ticker is stopped when the function exits
        defer ticker.Stop()

        for {
            select {
            case <-ticker.C:
                // Execute matchmaking logic on each tick
                m.FindMatches()
            }
        }
    }()
}