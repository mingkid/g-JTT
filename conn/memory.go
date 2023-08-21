package conn

import (
	"sync"
)

type MemoryConnPool struct {
	connections map[string]*Connection
	mu          sync.RWMutex // Mutex for concurrent access
}

func (cp *MemoryConnPool) Add(termID string, conn *Connection) {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	cp.connections[termID] = conn
}

func (cp *MemoryConnPool) Get(termID string) (c *Connection, ok bool) {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	cp.disconnectOnTimeout(termID)
	c, ok = cp.connections[termID]
	return
}

func (m *MemoryConnPool) disconnectOnTimeout(sn string) {
	if c, found := m.connections[sn]; found {
		if c.IsExpired() {
			delete(m.connections, sn)
		}
	}
}

var defaultConnPoolOnce sync.Once
var defaultConnPool Pool

func DefaultConnPool() Pool {
	defaultConnPoolOnce.Do(func() {
		defaultConnPool = &MemoryConnPool{connections: make(map[string]*Connection)}
	})
	return defaultConnPool
}
