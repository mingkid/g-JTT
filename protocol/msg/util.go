package msg

import "time"

const (
	TimeLayout    = "20060102150405" // 时间转换格式
	LimitlessTime = "000000000000"   // 无时间限制
)

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
