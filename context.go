package jtt

import (
	"net"

	"github.com/mingkid/g-jtt/conn"
	"github.com/mingkid/g-jtt/protocol/codec"
	"github.com/mingkid/g-jtt/protocol/msg"
)

type Context struct {
	c       *conn.Connection
	head    msg.Head
	rawData []byte
}

// Head 返回终端请求消息头
func (c *Context) Head() msg.Head {
	return c.head
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
	m := msg.M8001{
		Head: msg.Head{
			MsgID: msg.MsgIDPlatformCommResp,
			Phone: ctx.head.Phone,
		},
		AnswerSerialNo: ctx.head.SN,
		AnswerMsgID:    ctx.head.MsgID,
		Result:         res,
	}

	e := new(codec.Encoder)
	b, err := e.Encode(m)
	if err != nil {
		return err
	}
	return ctx.c.Send(b)
}
