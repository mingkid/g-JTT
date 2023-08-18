package conn

import (
	"sync"
)

type MemoryConnPool struct {
	connections map[string]*Connection
	mu          sync.RWMutex // Mutex for concurrent access
}

func (cp *MemoryConnPool) Add(terminalID string, conn *Connection) {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	cp.connections[terminalID] = conn
}

func (cp *MemoryConnPool) Get(terminalID string) *Connection {
	cp.mu.RLock()
	defer cp.mu.RUnlock()

	return cp.connections[terminalID]
}

var defaultConnPoolOnce sync.Once
var defaultConnPool Pool

func DefaultConnPool() Pool {
	defaultConnPoolOnce.Do(func() {
		defaultConnPool = &MemoryConnPool{connections: make(map[string]*Connection)}
	})
	return defaultConnPool
}
