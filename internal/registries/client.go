package registries

import (
	"net"
	"sync"

	"github.com/google/uuid"
	"github.com/leikyz/lkz-online-services/internal/models"
)

type ClientManager struct {
	clients map[string]*models.Client
	mu      sync.RWMutex
}

// Clients registry instance
var Clients = &ClientManager{clients: make(map[string]*models.Client)}

// Creates a new client and adds it to the registry
func (m *ClientManager) CreateClient(username string, level int, conn net.Conn) *models.Client {
	m.mu.Lock()
	defer m.mu.Unlock()

	c := &models.Client{
		ID:       uuid.New().String(),
		Username: username,
		Level:    level,
		Conn:     conn,
	}
	m.clients[c.ID] = c
	return c
}

// Remove removes a client from the internal map by ID
func (m *ClientManager) Remove(id string) {
	m.mu.Lock()         // Lock for write access
	defer m.mu.Unlock() // Unlock when the function exits

	// Check existence before deletion (optional but safe)
	if _, ok := m.clients[id]; ok {
		delete(m.clients, id)
	}
}

// Removes a client from the registry by ID
func (m *ClientManager) RemoveClient(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.clients, id)
}

// Retrieves a client by ID
func (m *ClientManager) GetByID(id string) (*models.Client, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	c, ok := m.clients[id]
	return c, ok
}
