package jtt

import (
	"net"
	"time"

	"gihub.com/mingkid/g-jtt/conn"
	"gihub.com/mingkid/g-jtt/protocol/codec"
	"gihub.com/mingkid/g-jtt/protocol/msg"
)

type Engine struct {
	connPool conn.Pool
	handlers map[msg.MsgID][]HandleFunc
}

func New(connPool conn.Pool) *Engine {
	return &Engine{
		connPool: connPool,
		handlers: make(map[msg.MsgID][]HandleFunc),
	}
}

func Default() *Engine {
	return New(conn.DefaultConnPool())
}

func (e *Engine) RegisterHandler(messageID msg.MsgID, handler HandleFunc) {
	e.handlers[messageID] = append(e.handlers[messageID], handler)
}

func (e *Engine) Serve(port string) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		c, err := listener.Accept()
		if err != nil {
			// Handle accept error
			continue
		}

		connection := conn.NewConnection(c, time.Now().Add(time.Minute)) // Adjust expiration time as needed

		go e.handleConnection(connection)
	}
}

func (e *Engine) handleConnection(c *conn.Connection) {
	rawData, _ := c.Receive() // Adjust error handling as needed

	ctx := &Context{
		conn:    c,
		rawData: rawData,
	}

	e.processMessage(ctx)
}

func (e *Engine) processMessage(ctx *Context) {
	messageID := codec.ExtraMsgID(ctx.Body()) // Extract message ID

	handlers, ok := e.handlers[messageID]
	if !ok {
		// Handle unknown message ID
		return
	}

	for _, handler := range handlers {
		handler(ctx)
	}
}

type HandleFunc func(ctx *Context)
