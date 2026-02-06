package models

import "sync"

type Lobby struct {
    ID      string           
    Clients []*Client         
    Mu      sync.RWMutex      
	
}

func (l *Lobby) IsFull() bool {
    l.Mu.RLock()
    defer l.Mu.RUnlock()
    
    return len(l.Clients) >= 4
}

func (l *Lobby) IsAvailable(c *Client) bool {
	l.Mu.RLock()
	defer l.Mu.RUnlock()

	for _, client := range l.Clients {
		if client == c || l.IsFull() {
			return false
		}
	}
	return true
}