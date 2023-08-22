package jtt

import (
	"net"
	"time"

	"github.com/mingkid/g-jtt/conn"
	"github.com/mingkid/g-jtt/protocol/codec"
	"github.com/mingkid/g-jtt/protocol/msg"
)

type Engine struct {
	PhoneToTermID PhoneToTermID
	connPool      conn.Pool
	handlers      map[msg.MsgID][]HandleFunc
}

func New(connPool conn.Pool) *Engine {
	return &Engine{
		connPool: connPool,
		handlers: make(map[msg.MsgID][]HandleFunc),
	}
}

func Default() (e *Engine) {
	e = New(conn.DefaultConnPool())
	e.PhoneToTermID = func(phone string) (termID string) {
		return phone
	}
	return
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
		rawConn, err := listener.Accept()
		if err != nil {
			// Handle accept error
			continue
		}

		c := conn.NewConnection(rawConn, time.Now().Add(time.Minute)) // Adjust expiration time as needed

		go func() {
			ctx := e.createContext(c)
			e.processMessage(ctx, c)
		}()
	}
}

func (e *Engine) createContext(c *conn.Connection) *Context {
	rawData, _ := c.Receive() // Adjust error handling as needed

	return &Context{
		c:       c,
		rawData: rawData,
	}
}

func (e *Engine) processMessage(ctx *Context, c *conn.Connection) {
	var (
		msgHead msg.Head
		decoder codec.Decoder
	)

	_ = decoder.Decode(msgHead, ctx.Data())

	ctx.head = msgHead

	// 更新连接池
	termID := e.PhoneToTermID(msgHead.Phone)
	if _, ok := e.connPool.Get(termID); !ok {
		e.connPool.Add(termID, c)
	}

	// 执行控制器函数
	handlers, ok := e.handlers[msgHead.MsgID]
	if !ok {
		// Handle unknown message ID
		return
	}
	for _, handler := range handlers {
		handler(ctx)
	}
}

type HandleFunc func(ctx *Context)
type PhoneToTermID func(phone string) (termID string)
