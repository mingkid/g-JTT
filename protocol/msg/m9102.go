package msg

// M9102 JT/T 1078 音视频实时传输控制
type M9102 struct {
	ChanNo              uint8                  // 逻辑通道号
	ControlDirective    M90102ControlDirective // 控制指令
	CloseAudioVideoType M9102CloseAVType       // 关闭音视频类型
	SwitchStreamType    M9102SwitchStreamType  // 切换码流类型
}

type M90102ControlDirective uint8

const (
	M9102ControlCloseAV      M90102ControlDirective = iota // 停止音视频码流
	M9102ControlSwitchStream                               // 切换码流
	M9102ControlPause                                      // 暂停
	M9102ControlResume                                     // 恢复
	M9102ControlCloseTalk                                  // 关闭双向对讲
)

type M9102CloseAVType uint8

const (
	M9102CloseAll   M9102CloseAVType = 0 // 关闭该通道有关的音视频数据
	M9102CloseAudio                  = 1 // 只关闭该通道有关的音频,保留该通道有关的视频
	M9102CloseVideo                  = 2 // 只关闭该通道有关的视频,保留该通道有关的音频
)

type M9102SwitchStreamType uint8

const (
	M9102SwitchToMainStream M9102SwitchStreamType = iota // 切换码流类型
	M9102SwitchToSubStream
)
