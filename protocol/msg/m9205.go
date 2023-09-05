package msg

import "time"

// M9205 数据结构，表示 JT/T 1078 协议中的9205数据
type M9205 struct {
	LogicChannelNumber uint8               // 逻辑通道号
	RawStartTime       string              `jtt13:"bcd,6"` // 开始时间，YYMMDDHHMMSS
	RawEndTime         string              `jtt13:"bcd,6"` // 结束时间，YYMMDDHHMMSS
	Warn               M9205Warn           // 报警标志
	AVResourceType     M9205AVResourceType // 音视频资源类型
	StreamType         M9205StreamType     // 码流类型
	StorageType        M9205StorageType    // 存储器类型
}

func NewM9205(chanNo uint8) *M9205 {
	return &M9205{
		LogicChannelNumber: chanNo,
		RawStartTime:       LimitlessTime,
		RawEndTime:         LimitlessTime,
		Warn:               0,
		AVResourceType:     0,
		StreamType:         0,
		StorageType:        0,
	}
}

// StartTime 返回开始时间条件
func (m *M9205) StartTime() time.Time {
	return ParseTime(m.RawStartTime)
}

// EndTime 返回结束时间条件
func (m *M9205) EndTime() time.Time {
	return ParseTime(m.RawEndTime)
}

// SetStartTime 设置开始时间条件
func (m *M9205) SetStartTime(t time.Time) {
	m.RawStartTime = t.Format(TimeLayout)[2:]
}

// SetEndTime 设置结束时间条件
func (m *M9205) SetEndTime(t time.Time) {
	m.RawEndTime = t.Format(TimeLayout)[2:]
}

// ClearStartTime 清除开始时间条件
func (m *M9205) ClearStartTime() {
	m.RawStartTime = LimitlessTime
}

// ClearEndTime 清除结束时间条件
func (m *M9205) ClearEndTime() {
	m.RawEndTime = LimitlessTime
}

// M9205AVResourceType 音视频资源类型
type M9205AVResourceType uint8

const (
	M9205AVResourceAV           M9205AVResourceType = iota // 音视频
	M9205AVResourceAudio                                   // 音频
	M9205AVResourceVideo                                   // 视频
	M9205AVResourceAudioOrVideo                            // 音频或视频
)

// M9205StreamType 码流类型
type M9205StreamType uint8

const (
	M9205StreamAll  M9205StreamType = iota // 所有码流
	M9205StreamMain                        // 主码流
	M9205StreamSub                         // 子码流
)

type M9205StorageType uint8

const (
	M9205StorageTypeAll    M9205StorageType = iota // 所有存储器
	M9205StorageTypeMain                           // 主存储器
	M9205StorageTypeBackup                         // 灾备存储器
)
