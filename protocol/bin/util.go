package bin

import (
	"encoding/binary"
	"gihub.com/mingkid/g-jtt/protocol/msg"
)

// ExtractMsgID 提取消息 ID
func ExtractMsgID(b []byte) msg.MsgID {
	return ExtractMsgIDFrom(b, 1)
}

// ExtractMsgIDFrom 从指定位置提取消息 ID
func ExtractMsgIDFrom(b []byte, start int) msg.MsgID {
	return msg.MsgID(binary.BigEndian.Uint16(b[start : start+1]))
}
