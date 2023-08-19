package msg

import (
	"errors"
	"fmt"
)

type MsgID uint16

const (
	// TermCommResp 终端通用应答
	TermCommResp MsgID = 0x0001

	// TermRegister 终端注册
	TermRegister = 0x0100

	// TermRegResp 终端注册应答
	TermRegResp = 0x8100

	// TermLocationRepose 终端位置汇报
	TermLocationRepose = 0x0200

	// TermLocationBatch 终端位置批量上传
	TermLocationBatch = 0x0704

	// TermAuth 终端鉴权
	TermAuth = 0x0102

	// PlatformCommResp 平台通用应答
	PlatformCommResp MsgID = 0x8001

	// TermHeartbeat 终端心跳
	TermHeartbeat MsgID = 0x0002
)

const MsgBodyMaxLength = 0x03ff

// Head is Message Head
type Head struct {
	MsgID MsgID
	BodyProps
	Version uint8  `jtt13:"-"`
	Phone   string `jtt13:"bcd,6" jtt19:"bcd,10"`
	SN      uint16
	//MsgPackagePacking
}

type BodyProps uint16

// Encrypted 加密返回 true，否而返回 false
func (prop BodyProps) Encrypted() bool {
	return (prop>>10)&0x01 == 0x001
}

// SetEncrypt 设置加密
func (prop BodyProps) SetEncrypt() {
	prop |= 0x01 << 10
}

// IsMultiplePackage 分包返回 true，否而返回 false
func (prop BodyProps) IsMultiplePackage() bool {
	return (prop>>13)&0x01 == 0x001
}

// SetMultiplePackage 设置分包
func (prop BodyProps) SetMultiplePackage() {
	prop |= 0x01 << 13
}

// BodyLength 消息体长度
func (prop BodyProps) BodyLength() uint16 {
	return uint16(prop & 0x03ff)
}

// SetBodyLength 设置消息体长度
func (prop BodyProps) SetBodyLength(len uint16) error {
	if len > MsgBodyMaxLength {
		return errors.New(fmt.Sprintf("消息体长度设置不能超过%v", MsgBodyMaxLength))
	}
	prop |= prop & BodyProps(len)
	return nil
}

type MsgPackagePacking struct {
	Total    uint16 `json:"total"`    // 消息包总数
	SerialNo uint16 `json:"serialNo"` // 包序号
}
