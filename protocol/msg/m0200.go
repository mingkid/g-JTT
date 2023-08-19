package msg

import "time"

// M0200 JT/T808 0200 数据包
type M0200 struct {
	Head
	Warn      M0200Warn
	Status    M0200Status
	Latitude  uint32
	Longitude uint32
	Altitude  uint16
	Speed     uint16
	Direction uint16
	Time      string `jtt13:"bcd,6"`
	Extra     M0200Extra
}

// LocateTime 定位时间
func (m M0200) LocateTime() (time.Time, error) {
	return time.ParseInLocation("20060102150405", "20"+m.Time, time.Local)
}
