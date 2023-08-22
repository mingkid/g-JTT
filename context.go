package jtt

import (
	"github.com/mingkid/g-jtt/conn"
	"github.com/mingkid/g-jtt/protocol/codec"
	"github.com/mingkid/g-jtt/protocol/msg"
)

type Context struct {
	c       *conn.Connection
	rawData []byte
}

func (ctx *Context) Data() []byte {
	return ctx.rawData
}

// RemoteAddr 返回远程网络地址（如果知道）
func (ctx *Context) RemoteAddr() Addr {
	return ctx.c.RemoteAddr()
}

func (ctx *Context) Generic(res msg.M8001Result) error {
	var (
		msg     msg.M8001
		encoder codec.Encoder
	)
	b, err := encoder.Encode(msg)
	if err != nil {
		return err
	}
	return ctx.c.Send(b)
}
