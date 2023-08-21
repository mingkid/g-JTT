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

// Receive 接收数据
func (c *Connection) Receive() ([]byte, error) {
	b := make([]byte, 1024)
	n, err := c.conn.Read(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}

// Send 发送数据
func (c *Connection) Send(b []byte) error {
	if _, err := c.conn.Write(b); err != nil {
		return err
	}
	return nil
}
