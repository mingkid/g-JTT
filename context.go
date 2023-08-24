package jtt

import (
	"net"

	"github.com/mingkid/g-jtt/conn"
	"github.com/mingkid/g-jtt/protocol/bin"
	"github.com/mingkid/g-jtt/protocol/codec"
	"github.com/mingkid/g-jtt/protocol/msg"
)

type Context struct {
	c       *conn.Connection // 连接
	head    msg.Head         // 消息头
	rawData []byte           // 终端原始数据
	termID  string           // 终端 ID
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
	m := msg.M8001{
		Head: msg.Head{
			MsgID: msg.MsgIDPlatformCommResp,
			Phone: ctx.head.Phone,
		},
		M8001Body: msg.M8001Body{
			AnswerSerialNo: ctx.head.SN,
			AnswerMsgID:    ctx.head.MsgID,
			Result:         res,
		},
	}

	e := new(codec.Encoder)
	b, err := e.Encode(m)
	if err = calcBodyLength(err, &m.Head, m.M8001Body); err != nil {
		return err
	}
	return ctx.c.Send(bin.Escape(b))
}

// Register 返回终端注册响应
func (ctx *Context) Register(res msg.M8100Result, token string) error {
	m := msg.M8100{
		Head: msg.Head{
			MsgID: msg.MsgIDTermRegResp,
			Phone: ctx.head.Phone,
		},
		M8100Body: msg.M8100Body{
			AnswerSerialNo: ctx.head.SN,
			Result:         res,
			Token:          token,
		},
	}

	e := new(codec.Encoder)
	b, err := e.Encode(m)
	if err = calcBodyLength(err, &m.Head, m.M8100Body); err != nil {
		return err
	}
	return ctx.c.Send(bin.Escape(b))
}

func calcBodyLength(err error, head *msg.Head, body any) error {
	size, err := bin.CalculateMsgLength(body)
	if err != nil {
		return err
	}
	if err = head.SetBodyLength(size); err != nil {
		return err
	}
	return nil
}
