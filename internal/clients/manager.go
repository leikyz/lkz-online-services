package client

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type Manager struct {
	clients map[string]*Client
	sync.RWMutex
}

var ClientManager = &Manager{
	clients: make(map[string]*Client),
}

func (m *Manager) CreateClient(username string, level int) *Client {
	m.Lock()
	defer m.Unlock()

	client := &Client{
		ID:       uuid.New().String(),
		Username: username,
		Level:    level,
	}

	client.ID = uuid.New().String()

	fmt.Printf("Creating client: %+v\n", client.ID)

	m.clients[client.ID] = client
	return client

}
