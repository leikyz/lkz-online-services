package registries

import (
	"net"
	"sync"
	"sync/atomic"

	"github.com/leikyz/lkz-online-services/internal/models"
)

type ClientManager struct {
	clients      map[uint32]*models.Client
	mu           sync.RWMutex
	lastClientID uint32
}

// Clients registry instance
var Clients = &ClientManager{clients: make(map[uint32]*models.Client)}

// Creates a new client and adds it to the registry
func (m *ClientManager) CreateClient(username string, level int, conn net.Conn) *models.Client {
	m.mu.Lock()
	defer m.mu.Unlock()

	newID := atomic.AddUint32(&m.lastClientID, 1)

	c := &models.Client{
		ID:       newID,
		Username: username,
		Level:    level,
		Conn:     conn,
	}
	m.clients[c.ID] = c
	return c
}

// Remove removes a client from the internal map by ID
func (m *ClientManager) Remove(id uint32) {
	m.mu.Lock()         // Lock for write access
	defer m.mu.Unlock() // Unlock when the function exits

	// Check existence before deletion (optional but safe)
	if _, ok := m.clients[id]; ok {
		delete(m.clients, id)
	}
}

// Removes a client from the registry by ID
func (m *ClientManager) RemoveClient(id uint32) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.clients, id)
}

// Retrieves a client by ID
func (m *ClientManager) GetByID(id uint32) (*models.Client, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	c, ok := m.clients[id]
	return c, ok
}
