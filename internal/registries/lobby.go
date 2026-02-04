package registries

import (
	"sync"

	"github.com/google/uuid"
	"github.com/leikyz/lkz-online-services/internal/models"
)

type LobbyManager struct {
	clients map[string]*models.Lobby
	mu      sync.RWMutex
}

// Clients registry instance
var Lobbies = &LobbyManager{clients: make(map[string]*models.Lobby)}

// Creates a new lobby and adds it to the registry
func (m *LobbyManager) CreateLobby(id int, maxClients int) *models.Lobby {
	m.mu.Lock()
	defer m.mu.Unlock()

	c := &models.Lobby{
		ID:      uuid.New().String(),
		Clients: make([]models.Client, 4), // Preallocate slice with capacity for 4 clients
	}
	m.clients[c.ID] = c
	return c
}

// Removes a lobby from the registry by ID
func (m *LobbyManager) RemoveLobby(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.clients, id)
}

// Retrieves a lobby by ID
func (m *LobbyManager) GetByID(id string) (*models.Lobby, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	c, ok := m.clients[id]
	return c, ok
}
