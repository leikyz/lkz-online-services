package registries

import (
	"sync"
	"sync/atomic"

	"github.com/leikyz/lkz-online-services/internal/models"
)

type LobbyManager struct {
	lobbies     map[uint32]*models.Lobby
	mu          sync.RWMutex
	lastLobbyId uint32
}

// Clients registry instance
var Lobbies = &LobbyManager{lobbies: make(map[uint32]*models.Lobby)}

// Creates a new lobby and adds it to the registry
func (m *LobbyManager) CreateLobby(id int, maxClients int) *models.Lobby {
	m.mu.Lock()
	defer m.mu.Unlock()

	newID := atomic.AddUint32(&m.lastLobbyId, 1)

	c := &models.Lobby{
		ID:      newID,
		Clients: make([]*models.Client, 0, 4),
	}

	m.lobbies[c.ID] = c
	return c
}

// Removes a lobby from the registry by ID
func (m *LobbyManager) RemoveLobby(id uint32) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.lobbies, id)
}

// Retrieves a lobby by ID
func (m *LobbyManager) GetByID(id uint32) (*models.Lobby, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	c, ok := m.lobbies[id]
	return c, ok
}
