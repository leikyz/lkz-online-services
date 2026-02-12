package registries

import (
	"sync"
	"sync/atomic"

	"github.com/leikyz/lkz-online-services/internal/models"
)

type SessionManager struct {
	sessions      map[uint32]*models.Session
	mu            sync.RWMutex
	lastSessionId uint32
}

// Clients registry instance
var Sessions = &SessionManager{sessions: make(map[uint32]*models.Session)}

// Creates a new Session and adds it to the registry
func (m *SessionManager) CreateSession(lobby *models.Lobby) *models.Session {
	m.mu.Lock()
	defer m.mu.Unlock()

	newID := atomic.AddUint32(&m.lastSessionId, 1)

	c := &models.Session{
		ID:    newID,
		Lobby: lobby,
	}

	m.sessions[c.ID] = c
	return c
}

// Removes a Session from the registry by ID
func (m *SessionManager) RemoveSession(id uint32) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.sessions, id)
}

// Retrieves a Session by ID
func (m *SessionManager) GetByID(id uint32) (*models.Session, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	c, ok := m.sessions[id]
	return c, ok
}
