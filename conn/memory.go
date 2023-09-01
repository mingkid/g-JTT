package conn

import (
	"sync"
)

type MemoryConnPool struct {
	connections map[string]*Connection
	mu          sync.RWMutex // Mutex for concurrent access
}

// Add 添加终端 ID 对应的连接对象
func (cp *MemoryConnPool) Add(termID string, conn *Connection) {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	cp.connections[termID] = conn
}

// Get 返回终端 ID 对应的连接对象
func (cp *MemoryConnPool) Get(termID string) (c *Connection, ok bool) {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	cp.disconnectOnTimeout(termID)
	c, ok = cp.connections[termID]
	return
}

// Remove 释放终端 ID 对应的连接对象
func (cp *MemoryConnPool) Remove(termID string) {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	if _, found := cp.connections[termID]; found {
		delete(cp.connections, termID)
	}
}

func (cp *MemoryConnPool) disconnectOnTimeout(termID string) {
	if c, found := cp.connections[termID]; found {
		if c.IsExpired() {
			delete(cp.connections, termID)
		}
	}
}

var defaultConnPoolOnce sync.Once
var defaultConnPool Pool

// DefaultConnPool 返回默认连接池
func DefaultConnPool() Pool {
	defaultConnPoolOnce.Do(func() {
		defaultConnPool = &MemoryConnPool{connections: make(map[string]*Connection)}
	})
	return defaultConnPool
}
