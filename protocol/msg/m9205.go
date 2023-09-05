package msg

// M9205 数据结构，表示 JT/T 1078 查询资源列表指令
type M9205 struct {
	LogicChannelNumber uint8               // 逻辑通道号
	Duration                               // 时间段条件
	Warn               M9205Warn           // 报警标志
	AVResourceType     M9205AVResourceType // 音视频资源类型
	StreamType         M9205StreamType     // 码流类型
	StorageType        M9205StorageType    // 存储器类型
}

// NewM9205 返回查询资源列表指令
func NewM9205(chanNo uint8) *M9205 {
	return &M9205{
		LogicChannelNumber: chanNo,
		Duration:           DurationMin(),
	}
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
