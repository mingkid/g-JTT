package jtt

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/mingkid/g-jtt/conn"
	"github.com/mingkid/g-jtt/protocol/bin"
	"github.com/mingkid/g-jtt/protocol/codec"
	"github.com/mingkid/g-jtt/protocol/msg"
)

// Context 请求上下文
type Context struct {
	c       *conn.Connection // 连接
	head    msg.Head         // 消息头
	rawData []byte           // 终端原始数据
	termID  string           // 终端 ID
}

// NewContext 创建上下文对象
func NewContext(c *conn.Connection) (*Context, error) {
	// 接收数据
	rawData, err := c.Receive()
	if err != nil {
		if err == io.EOF {
			return nil, err
		}
		return nil, errors.New(fmt.Sprintf("创建上下文异常: %s", err.Error()))
	}
	fmt.Printf("[JTT] %s | %s | 收到终端数据\n%s\n", time.Now().Format("2006/01/02 - 15:04:05"), c.RemoteAddr(), hex.EncodeToString(rawData))

	// 转义还原，验证校验码，删除校验位
	data := bin.Unescape(rawData)
	if err = bin.Verify(data[:len(data)-1], data[len(data)-1]); err != nil {
		return nil, errors.New(fmt.Sprintf("创建上下文异常: %s", err.Error()))
	}
	data = data[:len(data)-1]

	return &Context{
		c:       c,
		rawData: data,
	}, nil
}

// TermID 返回请求终端的 ID
func (ctx *Context) TermID() string {
	return ctx.termID
}

// Head 返回终端请求消息头
func (ctx *Context) Head() msg.Head {
	return ctx.head
}

// Data 返回终端发送的原始数据
func (ctx *Context) Data() []byte {
	return ctx.rawData
}

// RemoteAddr 返回远程网络地址（如果知道）
func (ctx *Context) RemoteAddr() net.Addr {
	return ctx.c.RemoteAddr()
}

// Generic 返回平台通用响应
func (ctx *Context) Generic(res msg.M8001Result) error {
	m := msg.Msg[msg.M8001]{
		Head: msg.Head{
			MsgID: msg.MsgIDPlatformCommResp,
			Phone: ctx.head.Phone,
		},
		Body: msg.M8001{
			AnswerSerialNo: ctx.head.SN,
			AnswerMsgID:    ctx.head.MsgID,
			Result:         res,
		},
	}

	b, err := packaging(&m)
	if err != nil {
		return err
	}

	err = ctx.c.Send(b)
	fmt.Printf("[JTT] %s | %s | 发送消息\n%s\n", time.Now().Format("2006/01/02 - 15:04:05"), ctx.RemoteAddr(), hex.EncodeToString(b))
	return err
}

// Register 返回终端注册响应
func (ctx *Context) Register(res msg.M8100Result, token string) error {
	m := msg.Msg[msg.M8100]{
		Head: msg.Head{
			MsgID: msg.MsgIDTermRegResp,
			Phone: ctx.head.Phone,
		},
		Body: msg.M8100{
			AnswerSerialNo: ctx.head.SN,
			Result:         res,
			Token:          token,
		},
	}

	b, err := packaging(&m)
	if err != nil {
		return err
	}

	err = ctx.c.Send(b)
	fmt.Printf("[JTT] %s | %s | 发送消息\n%s\n", time.Now().Format("2006/01/02 - 15:04:05"), ctx.RemoteAddr(), hex.EncodeToString(b))
	return err
}

// packaging 封装
func packaging[TBody any](m *msg.Msg[TBody]) ([]byte, error) {
	// 计算消息体长度
	if err := calcBodyLength(m); err != nil {
		return nil, err
	}

	// 编码
	e := new(codec.Encoder)
	b, err := e.Encode(m)
	if err != nil {
		return nil, err
	}

	// 填充校验码
	b = append(b, bin.Checksum(b))

	//  转义
	b = bin.Escape(b)
	return b, nil
}

// calcBodyLength 计算消息体长度
func calcBodyLength[TBody any](m *msg.Msg[TBody]) error {
	size, err := bin.CalculateMsgLength(m.Body)
	if err != nil {
		return err
	}
	if err = m.Head.SetBodyLength(size); err != nil {
		return err
	}
	return nil
}
