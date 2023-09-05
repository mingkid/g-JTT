package msg

import (
	"net"
)

// M9201 数据结构，表示 JT/T 1078 远程录像回放请求
type M9201 struct {
	VideoCtrl
	AVType      AVType      // 音视频类型
	StreamType  StreamType  // 码流类型
	StorageType StorageType // 存储器类型
	ReplayMode  ReplayMode  // 回放方式
	ReplaySpeed PlaySpeed   // 回放倍速
	Duration
}

func NewM9201(chanNo uint8) *M9201 {
	return &M9201{
		VideoCtrl: VideoCtrl{ChanNo: chanNo},
		Duration:  DurationMin(),
	}
}

// NewTCP9201 返回基于 TCP 通讯协议的远程录像回放请求指令
func NewTCP9201(ipAddr net.IP, port uint16, chanNo uint8) *M9201 {
	m := NewM9201(chanNo)
	m.SetTCPAddr(ipAddr, port)
	return m
}

// NewUDP9201 返回基于 UDP 通讯协议的远程录像回放请求指令
func NewUDP9201(ipAddr net.IP, port uint16, chanNo uint8) *M9201 {
	m := NewM9201(chanNo)
	m.SetUDPAddr(ipAddr, port)
	return m
}
