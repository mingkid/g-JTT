package msg

// M0200Status JT/T808消息类型0x0200的状态位定义
type M0200Status uint32

// ACCOn ACC开
func (s M0200Status) ACCOn() bool {
	return ((s >> 0) & 0x01) == 1
}

// PositionOn 定位开
func (s M0200Status) PositionOn() bool {
	return ((s >> 1) & 0x01) == 1
}

// IsSLat 南纬
func (s M0200Status) IsSLat() bool {
	return ((s >> 2) & 0x01) == 1
}

// IsWLong 西经
func (s M0200Status) IsWLong() bool {
	return ((s >> 3) & 0x01) == 1
}

// Operating 运营中
func (s M0200Status) Operating() bool {
	return ((s >> 4) & 0x01) != 1
}

// Encrypted 加密
func (s M0200Status) Encrypted() bool {
	return ((s >> 5) & 0x01) == 1
}

// EmergencyStop 紧急刹车
func (s M0200Status) EmergencyStop() bool {
	return ((s >> 6) & 0x01) == 1
}

// LaneDeparture 车道偏移
func (s M0200Status) LaneDeparture() bool {
	return ((s >> 7) & 0x01) == 1
}

// LoadStatus 荷载情况
func (s M0200Status) LoadStatus() LoadStatus {
	Is8thBit := (s>>8)&0x01 == 0
	Is9thBit := (s>>9)&0x01 == 0
	if !Is8thBit && Is9thBit {
		return LoadStatusHalf // 01 表示半载
	} else if Is8thBit && Is9thBit {
		return LoadStatusFully // 11 表示满载
	} else {
		return LoadStatusEmpty // 00 10 表示空车
	}
}

// OilChannelNormal 油路正常
func (s M0200Status) OilChannelNormal() bool {
	return ((s >> 10) & 0x01) != 1
}

// CircuitNormal 电路正常
func (s M0200Status) CircuitNormal() bool {
	return ((s >> 11) & 0x01) != 1
}

// DoorLocked 车门加锁
func (s M0200Status) DoorLocked() bool {
	return ((s >> 12) & 0x01) == 1
}

// FrontDoorOpened 前门打开
func (s M0200Status) FrontDoorOpened() bool {
	return ((s >> 13) & 0x01) == 1
}

// MidDoorOpened 中门打开
func (s M0200Status) MidDoorOpened() bool {
	return ((s >> 14) & 0x01) == 1
}

// BackDoorOpened 后门打开
func (s M0200Status) BackDoorOpened() bool {
	return ((s >> 15) & 0x01) == 1
}

// DriveRoomDoorOpened 驾驶席门打开
func (s M0200Status) DriveRoomDoorOpened() bool {
	return ((s >> 16) & 0x01) == 1
}

// ElseRoomDoorOpened 其它门打开
func (s M0200Status) ElseRoomDoorOpened() bool {
	return ((s >> 17) & 0x01) == 1
}

// GPSUsed 使用卫星定位
func (s M0200Status) GPSUsed() bool {
	return ((s >> 18) & 0x01) == 1
}

// BPSUsed 使用北斗卫星定位
func (s M0200Status) BPSUsed() bool {
	return ((s >> 19) & 0x01) == 1
}

// GLONASSUsed 使用GLONASS卫星定位
func (s M0200Status) GLONASSUsed() bool {
	return ((s >> 20) & 0x01) == 1
}

// GalileoUsed 使用Galileo卫星定位
func (s M0200Status) GalileoUsed() bool {
	return ((s >> 21) & 0x01) == 1
}

// TravelState 车辆处于行驶状态
func (s M0200Status) TravelState() bool {
	return ((s >> 22) & 0x01) == 1
}

// LoadStatus 荷载情况
type LoadStatus uint8

const (
	LoadStatusEmpty LoadStatus = 0x00
	LoadStatusHalf             = 0x01
	LoadStatusFully            = 0x11
)
