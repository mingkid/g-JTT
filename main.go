package jtt

import (
	"encoding/hex"
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
	PhoneToTermID        PhoneToTermID
	UnknownMsgHandleFunc HandleFunc
	connPool             conn.Pool
	handlers             map[msg.MsgID][]HandleFunc
}

func New(connPool conn.Pool) *Engine {
	return &Engine{
		connPool: connPool,
		handlers: make(map[msg.MsgID][]HandleFunc),
	}
}

func Default() (e *Engine) {
	e = New(conn.DefaultConnPool())
	e.PhoneToTermID = DefaultPhoneToTermID
	e.UnknownMsgHandleFunc = DefaultUnknownMsgHandle
	return
}

func (e *Engine) RegisterHandler(messageID msg.MsgID, handler HandleFunc) {
	e.handlers[messageID] = append(e.handlers[messageID], handler)
}

func (e *Engine) Serve(ip string, port uint) error {
	if err := e.checkServeRequirement(); err != nil {
		return err
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return err
	}
	defer listener.Close()
	fmt.Printf("[JTT] %s | 监听开始: %s! \n", time.Now().Format("2006/01/02 - 15:04:05"), listener.Addr())

	for {
		// 创建连接对象
		rawConn, err := listener.Accept()
		if err != nil {
			fmt.Printf("[JTT] %s | 终端 %s 连接异常 \n%s\n", time.Now().Format("2006/01/02 - 15:04:05"), rawConn.RemoteAddr(), err.Error())
			continue
		}
		fmt.Printf("[JTT] %s | 终端 %s 已连接！ \n", time.Now().Format("2006/01/02 - 15:04:05"), rawConn.RemoteAddr())
		c := conn.NewConnection(rawConn, time.Now().Add(time.Minute))

		go func() {
			for {
				// 创建上下文对象
				ctx, err := e.createContext(c)
				if err != nil {
					if err == io.EOF {
						fmt.Printf("[JTT] %s | 终端 %s 已断开连接！ \n", time.Now().Format("2006/01/02 - 15:04:05"), c.RemoteAddr())
						break
					}
					fmt.Printf("[JTT] %s | %s", time.Now().Format("2006/01/02 - 15:04:05"), err.Error())
					continue
				}

				// 终端消息处理
				if err = e.processMessage(ctx); err != nil {
					fmt.Printf("[JTT] %s | %s", time.Now().Format("2006/01/02 - 15:04:05"), err.Error())
					continue
				}

				// 连接添加到连接池
				e.connPoolAppend(ctx.termID, c)
			}
		}()
	}
}

func (e *Engine) createContext(c *conn.Connection) (*Context, error) {
	rawData, err := c.Receive()
	if err != nil {
		if err == io.EOF {
			return nil, err
		}
		return nil, errors.New(fmt.Sprintf("%s Creating Context Error: %s", c.RemoteAddr(), err.Error()))
	}

	return &Context{
		c:       c,
		rawData: rawData,
	}, nil
}

func (e *Engine) processMessage(ctx *Context) (err error) {
	var (
		msgHead msg.Head
		decoder codec.Decoder
	)

	_ = decoder.Decode(&msgHead, ctx.Data())

	// 补充上下文信息
	ctx.head = msgHead
	if ctx.termID, err = e.PhoneToTermID(msgHead.Phone); err != nil {
		return errors.New(fmt.Sprintf("Phone %s to TermID Error: %s", msgHead.Phone, err.Error()))
	}

	// 执行控制器函数
	handlers, ok := e.handlers[msgHead.MsgID]
	if !ok {
		e.UnknownMsgHandleFunc(ctx)
	}
	for _, handler := range handlers {
		handler(ctx)
	}

	return
}

func (e *Engine) connPoolAppend(termID string, c *conn.Connection) {
	if _, ok := e.connPool.Get(termID); !ok {
		e.connPool.Add(termID, c)
	}
}

func (e *Engine) checkServeRequirement() error {
	if e.PhoneToTermID == nil {
		return errors.New("PhoneToTermID method is not implemented")
	}
	if e.UnknownMsgHandleFunc == nil {
		return errors.New("UnknownMsgHandleFunc method is not implemented")
	}
	return nil
}

type HandleFunc func(ctx *Context)
type PhoneToTermID func(phone string) (termID string, err error)

// DefaultPhoneToTermID 默认手机号码转终端 ID 控制器函数
var DefaultPhoneToTermID = func(phone string) (termID string, err error) {
	if phone == "" {
		return "", errors.New("TermID Not Found")
	}
	return phone, nil
}

// DefaultUnknownMsgHandle 默认未知消息处理控制器函数
var DefaultUnknownMsgHandle = func(ctx *Context) {
	fmt.Printf("[JTT] %s | %s | 未知消息处理\n%s\n", time.Now().Format("2006/01/02 - 15:04:05"), ctx.RemoteAddr(), hex.EncodeToString(ctx.Data()))
	if ctx.Head().MsgID == msg.MsgIDTermRegister {
		_ = ctx.Register(msg.M8100ResultTermRegistered, "")
	} else {
		_ = ctx.Generic(msg.M8001ResultFail)
	}
}
