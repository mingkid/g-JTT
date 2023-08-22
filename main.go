package jtt

import (
	"errors"
	"fmt"
	"io"
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
	fmt.Printf("[JTT] 监听开始: %s! \n", listener.Addr())

	for {
		rawConn, err := listener.Accept()
		if err != nil {
			// Handle accept error
			continue
		}
		fmt.Printf("[JTT] 终端 %s 已连接！ \n", rawConn.RemoteAddr())

		c := conn.NewConnection(rawConn, time.Now().Add(time.Minute))

		go func() {
			for {
				ctx, err := e.createContext(c)
				if err != nil {
					fmt.Printf("[JTT] %s", err.Error())
				}

				e.processMessage(ctx)
				e.connPoolAppend(ctx.termID, c)
			}
		}()
	}
}

func (e *Engine) createContext(c *conn.Connection) (*Context, error) {
	rawData, err := c.Receive()
	if err != nil {
		if err == io.EOF {
			return nil, errors.New(fmt.Sprintf("终端 %s 已断开连接！ \n", c.RemoteAddr()))
		}
		fmt.Println(err)
	}

	return &Context{
		c:       c,
		rawData: rawData,
	}, nil
}

func (e *Engine) processMessage(ctx *Context) {
	var (
		msgHead msg.Head
		decoder codec.Decoder
	)

	_ = decoder.Decode(msgHead, ctx.Data())

	// 补充上下文信息
	ctx.head = msgHead
	ctx.termID = e.PhoneToTermID(msgHead.Phone)

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

func (e *Engine) connPoolAppend(termID string, c *conn.Connection) {
	if _, ok := e.connPool.Get(termID); !ok {
		e.connPool.Add(termID, c)
	}
}

type HandleFunc func(ctx *Context)
type PhoneToTermID func(phone string) (termID string)
