package msg

import "time"

const (
	TimeLayout    = "20060102150405" // 时间转换格式
	LimitlessTime = "000000000000"   // 无时间限制
)

// Duration 时间段控制
type Duration struct {
	RawStartTime string `jtt13:"bcd,6"` // 开始时间，YYMMDDHHMMSS
	RawEndTime   string `jtt13:"bcd,6"` // 结束时间，YYMMDDHHMMSS
}

// DurationMin 最小时间段
func DurationMin() Duration {
	return Duration{
		RawStartTime: LimitlessTime,
		RawEndTime:   LimitlessTime,
	}
}

func NewDuration(start, end time.Time) Duration {
	d := DurationMin()
	d.SetStartTime(start)
	d.SetEndTime(end)
	return d
}

// StartTime 返回开始时间条件
func (m *Duration) StartTime() time.Time {
	return ParseTime(m.RawStartTime)
}

// EndTime 返回结束时间条件
func (m *Duration) EndTime() time.Time {
	return ParseTime(m.RawEndTime)
}

// SetStartTime 设置开始时间条件
func (m *Duration) SetStartTime(t time.Time) {
	m.RawStartTime = t.Format(TimeLayout)[2:]
}

// SetEndTime 设置结束时间条件
func (m *Duration) SetEndTime(t time.Time) {
	m.RawEndTime = t.Format(TimeLayout)[2:]
}

// ClearStartTime 清除开始时间条件
func (m *Duration) ClearStartTime() {
	m.RawStartTime = LimitlessTime
}

// ClearEndTime 清除结束时间条件
func (m *Duration) ClearEndTime() {
	m.RawEndTime = LimitlessTime
}

// ParseTime 返回 BCD 时间转换，t 格式为 YYMMDDHHMMSS
func ParseTime(t string) time.Time {
	if t == LimitlessTime || t == "" {
		return time.Time{}
	}
	if res, err := time.Parse(TimeLayout, "20"+t); err != nil {
		return time.Time{}
	} else {
		return res
	}
}
