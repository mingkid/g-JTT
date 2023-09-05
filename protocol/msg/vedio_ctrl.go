package msg

import "net"

// VideoCtrl 视频控制指令
type VideoCtrl struct {
	IPAddrLen uint8  // 服务器IP 地址长度
	IPAddr    string // 服务器IP 地址
	TCPPort   uint16 // 服务器视频通道监听端口号(TCP)
	UDPPort   uint16 // 服务器视频通道监听端口号(UDP)
	ChanNo    uint8  // 逻辑通道号
}

// SetServerIPAddr 设置服务器地址
func (m *VideoCtrl) SetServerIPAddr(addr net.IP) {
	m.IPAddr = addr.String()
	m.IPAddrLen = uint8(len(m.IPAddr))
}

// SetTCPAddr 设置 TCP 服务器地址
func (m *VideoCtrl) SetTCPAddr(addr net.IP, port uint16) {
	m.SetServerIPAddr(addr)
	m.TCPPort = port
}

// SetUDPAddr 设置 UDP 服务器地址
func (m *VideoCtrl) SetUDPAddr(addr net.IP, port uint16) {
	m.SetServerIPAddr(addr)
	m.UDPPort = port
}

// AVType 音视频资源类型
type AVType uint8

const (
	AVTypeAV           AVType = iota // 音视频
	AVTypeAudio                      // 音频
	AVTypeVideo                      // 视频
	AVTypeAudioOrVideo               // 音频或视频
)

// StreamType 码流类型
type StreamType uint8

const (
	StreamTypeAll  StreamType = iota // 所有码流
	StreamTypeMain                   // 主码流
	StreamTypeSub                    // 子码流
)

// StorageType 存储类型
type StorageType uint8

const (
	StorageTypeAll    StorageType = iota // 所有存储器
	StorageTypeMain                      // 主存储器
	StorageTypeBackup                    // 灾备存储器
)
