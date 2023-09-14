package jtt

import (
	"github.com/mingkid/g-jtt/protocol/bin"
	"github.com/mingkid/g-jtt/protocol/codec"
	"github.com/mingkid/g-jtt/protocol/msg"
)

// Send 下发指令
func Send[TBody any](e *Engine, m msg.Msg[TBody]) error {
	b, err := packaging(&m)
	if err != nil {
		return err
	}

	return SendBytes(e, m.Head.Phone, b)
}

// SendBytes 下发二进制数据指令
func SendBytes(e *Engine, phone string, b []byte) error {
	termID, err := e.PhoneToTermID(phone)
	if err != nil {
		return err
	}

	c, hasConn := e.connPool.Get(termID)
	if !hasConn {
		return newDeviceoffline(termID)
	}

	return c.Send(b)
}

// packaging 封装
func packaging[TBody any](m *msg.Msg[TBody]) ([]byte, error) {
	// 计算消息体长度
	if err := calcBodyLength(&m.Head, m.Body); err != nil {
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
func calcBodyLength(h *msg.Head, b any) error {
	size, err := bin.CalculateMsgLength(b)
	if err != nil {
		return err
	}
	if err = h.SetBodyLength(size); err != nil {
		return err
	}
	return nil
}
