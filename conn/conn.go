package conn

import (
	"errors"
	"net"
	"sync"
	"time"

	"github.com/mingkid/g-jtt/protocol/msg"
)

type Connection struct {
	conn       net.Conn
	status     ConnStatus
	expiration time.Time
	mu         sync.Mutex // Mutex for concurrent access
}

func NewConnection(conn net.Conn, expiration time.Time) *Connection {
	return &Connection{
		conn:       conn,
		expiration: expiration,
	}
}

// Status 连接状态
func (c *Connection) Status() ConnStatus {
	return c.status
}

// IsExpired 返回是否过期
func (c *Connection) IsExpired() bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	return time.Now().After(c.expiration)
}

// SetExpiration 设置到期时间
func (c *Connection) SetExpiration(expiration time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.expiration = expiration
}

// SetExpirationByDuration 设置连接还有多久就过期
func (c *Connection) SetExpirationByDuration(duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.expiration = time.Now().Add(duration)
}

// SetExpirationByTimestamp 设置到期时间戳
func (c *Connection) SetExpirationByTimestamp(timestamp int64) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.expiration = time.Unix(timestamp, 0)
}

// Receive 接收数据
func (c *Connection) Receive() ([]byte, error) {
	if c.IsExpired() {
		return nil, errors.New("连接已过期")
	}
	b := make([]byte, 1024)
	n, err := c.conn.Read(b)
	return b[:n], err
}

// Send 发送数据
func (c *Connection) Send(b []byte) error {
	if c.IsExpired() {
		return errors.New("连接已过期")
	}
	if _, err := c.conn.Write(b); err != nil {
		return err
	}
	return nil
}

// Close 关闭连接
func (c *Connection) Close() error {
	c.status = ConnStatusUnConnected
	return c.conn.Close()
}

// RemoteAddr 返回远程网络地址（如果知道）
func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

// Register 终端注册
func (c *Connection) Register(res msg.M8100Result) {
	if res == msg.M8100ResultSuccess {
		c.status = ConnStatusToAuth
	} else {
		c.status = ConnStatusToRegister
	}
}

// Auth 终端鉴权
func (c *Connection) Auth(res msg.M8001Result) {
	if res == msg.M8001ResultSuccess {
		c.status = ConnStatusOnline
	} else {
		c.status = ConnStatusToAuth
	}
}

// ConnStatus 终端状态
type ConnStatus uint8

const (
	ConnStatusUnConnected ConnStatus = iota // 未连接
	ConnStatusToRegister                    // 注册中
	ConnStatusToAuth                        // 鉴权中
	ConnStatusOnline                        // 在线
)
