package client

import (
	"sync"

	"github.com/google/uuid"
)

type Manager struct {
	clients map[string]*Client
	mu      sync.RWMutex
}

var ClientManager = &Manager{
	clients: make(map[string]*Client),
}

func (m *Manager) CreateClient(username string, level int) *Client {
	m.mu.Lock()
	defer m.mu.Unlock()

	c := &Client{
		ID:       uuid.New().String(),
		Username: username,
		Level:    level,
	}
	m.clients[c.ID] = c
	return c
}

func (m *Manager) RemoveClient(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.clients, id)
}

func (m *Manager) GetByID(id string) (*Client, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	c, ok := m.clients[id]
	return c, ok
}
