package registries

import (
	"sync"

	"github.com/google/uuid"
	"github.com/leikyz/lkz-online-services/internal/models"
)

type Manager struct {
	clients map[string]*models.Client
	mu      sync.RWMutex
}

// Clients registry instance
var Clients = &Manager{clients: make(map[string]*models.Client)}

// Creates a new client and adds it to the registry
func (m *Manager) CreateClient(username string, level int) *models.Client {
	m.mu.Lock()
	defer m.mu.Unlock()

	c := &models.Client{
		ID:       uuid.New().String(),
		Username: username,
		Level:    level,
	}
	m.clients[c.ID] = c
	return c
}

// Removes a client from the registry by ID
func (m *Manager) RemoveClient(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.clients, id)
}

// Retrieves a client by ID
func (m *Manager) GetByID(id string) (*models.Client, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	c, ok := m.clients[id]
	return c, ok
}
