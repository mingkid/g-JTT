package jtt

import (
	"gihub.com/mingkid/g-jtt/conn"
	"gihub.com/mingkid/g-jtt/protocol/codec"
	"gihub.com/mingkid/g-jtt/protocol/msg"
)

type Context struct {
	conn    *conn.Connection
	rawData []byte
}

func (ctx *Context) Body() []byte {
	return ctx.rawData
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
	return ctx.conn.Send(b)
}
