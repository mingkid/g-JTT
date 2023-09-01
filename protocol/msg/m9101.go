package msg

import (
	"net"
)

// M9101 JT/T 1078 9101数据
type M9101 struct {
	ServerIPAddressLength uint8           // 服务器IP 地址长度
	ServerIPAddress       string          // 服务器IP 地址
	TCPListenPort         uint16          // 服务器视频通道监听端口号(TCP)
	UDPListenPort         uint16          // 服务器视频通道监听端口号(UDP)
	LogicalChannel        uint8           // 逻辑通道号
	DataType              M9101DataType   // 数据类型
	StreamType            M9101StreamType // 码流类型
}

// SetTCPAddr 设置 TCP 服务器地址
func (m *M9101) SetTCPAddr(addr net.IP, port uint16) {
	m.SetServerIPAddr(addr)
	m.TCPListenPort = port
}

// SetUDPAddr 设置 UDP 服务器地址
func (m *M9101) SetUDPAddr(addr net.IP, port uint16) {
	m.SetServerIPAddr(addr)
	m.UDPListenPort = port
}

// SetServerIPAddr 设置服务器 IP 地址
func (m *M9101) SetServerIPAddr(addr net.IP) {
	m.ServerIPAddress = addr.String()
	m.ServerIPAddressLength = uint8(len(m.ServerIPAddress))
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
