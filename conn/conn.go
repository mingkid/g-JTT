package conn

import (
	"net"
	"sync"
	"time"
)

type Connection struct {
	conn       net.Conn
	expiration time.Time
	mu         sync.Mutex // Mutex for concurrent access
}

func NewConnection(conn net.Conn, expiration time.Time) *Connection {
	return &Connection{
		conn:       conn,
		expiration: expiration,
	}
}

func (c *Connection) IsExpired() bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	return time.Now().After(c.expiration)
}

func (c *Connection) SetExpiration(expiration time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.expiration = expiration
}

func (c *Connection) SetExpirationByDuration(duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.expiration = time.Now().Add(duration)
}

func (c *Connection) SetExpirationByTimestamp(timestamp int64) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.expiration = time.Unix(timestamp, 0)
}

func (c *Connection) Receive() ([]byte, error) {
	return nil, nil
}

func (c *Connection) Send(data []byte) error {
	return nil
}
