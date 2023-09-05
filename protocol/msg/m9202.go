package msg

import "time"

type M9202 struct {
	ChanNo      uint8         // 通道号
	ReplayCtrl  M9202PlayCtrl // 回放控制
	ReplaySpeed PlaySpeed     // 回放速度
	ReplayTime  string        `jtt13:"bcd,6"` // 回放时间
}

func NewM9202(chanNo uint8) *M9202 {
	return &M9202{
		ChanNo:     chanNo,
		ReplayTime: LimitlessTime,
	}
}

// SetReplayTime  设置回放时间，并将回放控制设置为拖动回放
func (m *M9202) SetReplayTime(t time.Time) {
	m.ReplayTime = t.Format(TimeLayout)[2:]
	m.ReplayCtrl = M9202PlayDragRewind
}

// CleanReplayTime  清除回放时间，并将回放控制设置为开始播放
func (m *M9202) CleanReplayTime(t time.Time) {
	m.ReplayTime = LimitlessTime
	m.ReplayCtrl = M9202PlayStart
}

type M9202PlayCtrl uint8

const (
	M9202PlayStart              M9202PlayCtrl = iota // 开始播放
	M9202PlayPause                                   // 暂停播放
	M9202PlayStop                                    // 结束播放
	M9202PlayFastForward                             // 快进
	M9202PlayKeyFrameFastRewind                      // 关键帧快退
	M9202PlayDragRewind                              // 拖回
	M9202PlayKeyFramePlay                            // 关键帧播放
)
