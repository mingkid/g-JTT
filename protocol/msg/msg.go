package msg

import (
	"errors"
	"fmt"
)

type MsgID uint16

func (id MsgID) Hex() string {
	return fmt.Sprintf("%04x", id)
}

const (
	MsgIDTermCommResp       MsgID = 0x0001 // 终端通用应答
	MsgIDTermRegister       MsgID = 0x0100 // 终端注册
	MsgIDTermRegResp        MsgID = 0x8100 // 终端注册应答
	MsgIDTermLocationReport MsgID = 0x0200 // 终端位置汇报
	MsgIDTermLocationBatch  MsgID = 0x0704 // 终端位置批量上传
	MsgIDTermAuth           MsgID = 0x0102 // 终端鉴权
	MsgIDPlatformCommResp   MsgID = 0x8001 // 平台通用应答
	MsgIDTermHeartbeat      MsgID = 0x0002 // 终端心跳
	MsgIDRealtimePlay       MsgID = 0x9101 // 实时音视频传输请求
	MsgIDRealtimePlayCtrl   MsgID = 0x9102 // 实时音视频传输控制
	MsgIDRealtimePlayStatus MsgID = 0x9105 // 实时音视频传输状态
	MsgIDPlayback           MsgID = 0x9201 // 远程录像回放
	MsgIDPlaybackCtrl       MsgID = 0x9202 // 远程录像回放控制
	MsgIDPlaybackList       MsgID = 0x9205 // 查询远程录像资源列表
)

const MsgBodyMaxLength = 0x03ff

// Msg 消息
type Msg[TBody any] struct {
	Head
	Body TBody
}

// Head 消息头
type Head struct {
	MsgID MsgID
	BodyProps
	Version uint8  `jtt13:"-"`
	Phone   string `jtt13:"bcd,6" jtt19:"bcd,10"`
	SN      uint16
	//MsgPackagePacking
}

// BodyProps 消息体属性
type BodyProps uint16

// Encrypted 加密返回 true，否而返回 false
func (prop BodyProps) Encrypted() bool {
	return (prop>>10)&0x01 == 0x001
}

// SetEncrypt 设置加密
func (prop *BodyProps) SetEncrypt() {
	*prop |= 0x01 << 10
}

// IsMultiplePackage 分包返回 true，否而返回 false
func (prop BodyProps) IsMultiplePackage() bool {
	return (prop>>13)&0x01 == 0x001
}

// SetMultiplePackage 设置分包
func (prop *BodyProps) SetMultiplePackage() {
	*prop |= 0x01 << 13
}

// BodyLength 消息体长度
func (prop BodyProps) BodyLength() uint16 {
	return uint16(prop & 0x03ff)
}

// SetBodyLength 设置消息体长度
func (prop *BodyProps) SetBodyLength(len uint16) error {
	if len > MsgBodyMaxLength {
		return errors.New(fmt.Sprintf("消息体长度设置不能超过%v", MsgBodyMaxLength))
	}
	*prop |= (*prop & 0xfc00) | BodyProps(len)
	return nil
}

// MsgPackagePacking 分包项
type MsgPackagePacking struct {
	Total    uint16 `json:"total"`    // 消息包总数
	SerialNo uint16 `json:"serialNo"` // 包序号
}
