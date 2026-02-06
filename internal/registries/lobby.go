package registries

import (
	"sync"

	"github.com/google/uuid"
	"github.com/leikyz/lkz-online-services/internal/models"
)

type LobbyManager struct {
	lobbies map[string]*models.Lobby
	mu      sync.RWMutex
}

// Clients registry instance
var Lobbies = &LobbyManager{lobbies: make(map[string]*models.Lobby)}

// Creates a new lobby and adds it to the registry
func (m *LobbyManager) CreateLobby(id int, maxClients int) *models.Lobby {
    m.mu.Lock()
    defer m.mu.Unlock()

    c := &models.Lobby{
        ID:      uuid.New().String(),
        Clients: make([]*models.Client, 0, 4), 
    }
    
    m.lobbies[c.ID] = c
    return c
}

// Removes a lobby from the registry by ID
func (m *LobbyManager) RemoveLobby(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.lobbies, id)
}

// Retrieves a lobby by ID
func (m *LobbyManager) GetByID(id string) (*models.Lobby, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	c, ok := m.lobbies[id]
	return c, ok
}
