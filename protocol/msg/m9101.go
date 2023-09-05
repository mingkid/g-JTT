package msg

import "net"

// M9101 JT/T 1078 9101数据
type M9101 struct {
	VideoCtrl
	DataType   M9101DataType   // 数据类型
	StreamType M9101StreamType // 码流类型
}

func NewM9101(chanNo uint8) *M9101 {
	m := &M9101{
		VideoCtrl: VideoCtrl{
			ChanNo: chanNo,
		},
	}
	return m
}

// NewTCP9101 返回基于 TCP 通讯协议的 0x9101 指令
func NewTCP9101(ipAddr net.IP, port uint16, chanNo uint8) *M9101 {
	m := NewM9101(chanNo)
	m.SetTCPAddr(ipAddr, port)
	return m
}

// NewUDP9101 返回基于 UDP 通讯协议的 0x9101 指令
func NewUDP9101(ipAddr net.IP, port uint16, chanNo uint8) *M9101 {
	m := NewM9101(chanNo)
	m.SetUDPAddr(ipAddr, port)
	return m
}

// M9101DataType 数据类型
type M9101DataType uint8

const (
	// M9101DataTypeAudioVideo 音视频
	M9101DataTypeAudioVideo M9101DataType = iota

	// M9101DataTypeVideo 视频
	M9101DataTypeVideo

	// M9101DataTypeTwoWayIntercom 双向对讲
	M9101DataTypeTwoWayIntercom

	// M9101DataTypeMonitor 监听
	M9101DataTypeMonitor

	// M9101DataTypeCenterBroadcast 中心广播
	M9101DataTypeCenterBroadcast

	// M9101DataTypePassThrough 透传
	M9101DataTypePassThrough
)

// M9101StreamType 码流类型
type M9101StreamType byte

const (
	// M9101StreamTypeMain 主码流
	M9101StreamTypeMain M9101StreamType = iota

	// M9101StreamTypeSub 子码流
	M9101StreamTypeSub
)
