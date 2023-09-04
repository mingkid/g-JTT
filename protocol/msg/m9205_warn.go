package msg

type M9205Warn uint64

// Emergency 紧急报警
func (w M9205Warn) Emergency() bool {
	return ((w >> 0) & 0x01) == 1
}

func (w *M9205Warn) SetEmergency(value bool) {
	if value {
		*w |= 1 << 0
	} else {
		*w &^= 1 << 0
	}
}

// OverSpeed 超速
func (w M9205Warn) OverSpeed() bool {
	return ((w >> 1) & 0x01) == 1
}

func (w *M9205Warn) SetOverSpeed(value bool) {
	if value {
		*w |= 1 << 1
	} else {
		*w &^= 1 << 1
	}
}

// Tired 疲劳驾驶
func (w M9205Warn) Tired() bool {
	return ((w >> 2) & 0x01) == 1
}

func (w *M9205Warn) SetTired(value bool) {
	if value {
		*w |= 1 << 2
	} else {
		*w &^= 1 << 2
	}
}

// Danger 危险驾驶行为
func (w M9205Warn) Danger() bool {
	return ((w >> 3) & 0x01) == 1
}

func (w *M9205Warn) SetDanger(value bool) {
	if value {
		*w |= 1 << 3
	} else {
		*w &^= 1 << 3
	}
}

// TermFault GNSS模块驾驶行为
func (w M9205Warn) TermFault() bool {
	return ((w >> 4) & 0x01) == 1
}

func (w *M9205Warn) SetTermFault(value bool) {
	if value {
		*w |= 1 << 4
	} else {
		*w &^= 1 << 4
	}
}

// AerialUnConn GNSS天线未接/被剪断
func (w M9205Warn) AerialUnConn() bool {
	return ((w >> 5) & 0x01) == 1
}

func (w *M9205Warn) SetAerialUnConn(b bool) {
	if b {
		*w |= 1 << 5
	} else {
		*w &^= 1 << 5
	}
}

// AerialShortCircuit GNSS天线短路
func (w M9205Warn) AerialShortCircuit() bool {
	return ((w >> 6) & 0x01) == 1
}

func (w *M9205Warn) SetAerialShortCircuit(b bool) {
	if b {
		*w |= 1 << 6
	} else {
		*w &^= 1 << 6
	}
}

// TermUndervoltage 终端主电源欠压
func (w M9205Warn) TermUndervoltage() bool {
	return ((w >> 7) & 0x01) == 1
}

func (w *M9205Warn) SetTermUndervoltage(b bool) {
	if b {
		*w |= 1 << 7
	} else {
		*w &^= 1 << 7
	}
}

// TermPowerFail 终端主电源掉电
func (w M9205Warn) TermPowerFail() bool {
	return ((w >> 8) & 0x01) == 1
}

func (w *M9205Warn) SetTermPowerFail(b bool) {
	if b {
		*w |= 1 << 8
	} else {
		*w &^= 1 << 8
	}
}

// LCDFault 终端LCD或显示器故障
func (w M9205Warn) LCDFault() bool {
	return ((w >> 9) & 0x01) == 1
}

func (w *M9205Warn) SetLCDFault(b bool) {
	if b {
		*w |= 1 << 9
	} else {
		*w &^= 1 << 9
	}
}

// TTSFault TTS模块故障
func (w M9205Warn) TTSFault() bool {
	return ((w >> 10) & 0x01) == 1
}

func (w *M9205Warn) SetTTSFault(b bool) {
	if b {
		*w |= 1 << 10
	} else {
		*w &^= 1 << 10
	}
}

// CameraFault 摄像头故障
func (w M9205Warn) CameraFault() bool {
	return ((w >> 11) & 0x01) == 1
}

func (w *M9205Warn) SetCameraFault(b bool) {
	if b {
		*w |= 1 << 11
	} else {
		*w &^= 1 << 11
	}
}

// TotalDriveTimeout 累计驾驶时长超时
func (w M9205Warn) TotalDriveTimeout() bool {
	return ((w >> 18) & 0x01) == 1
}

func (w *M9205Warn) SetTotalDriveTimeout(b bool) {
	if b {
		*w |= 1 << 18
	} else {
		*w &^= 1 << 18
	}
}

// StopTimeout 停车超时
func (w M9205Warn) StopTimeout() bool {
	return ((w >> 19) & 0x01) == 1
}

func (w *M9205Warn) SetStopTimeout(b bool) {
	if b {
		*w |= 1 << 19
	} else {
		*w &^= 1 << 19
	}
}

// AreaIO 区域进出
func (w M9205Warn) AreaIO() bool {
	return ((w >> 20) & 0x01) == 1
}

func (w *M9205Warn) SetAreaIO(b bool) {
	if b {
		*w |= 1 << 20
	} else {
		*w &^= 1 << 20
	}
}

// RoadIO 路线进出
func (w M9205Warn) RoadIO() bool {
	return ((w >> 21) & 0x01) == 1
}

func (w *M9205Warn) SetRoadIO(b bool) {
	if b {
		*w |= 1 << 21
	} else {
		*w &^= 1 << 21
	}
}

// RoadDriveNotEnoughTime 路线驾驶时长不足
func (w M9205Warn) RoadDriveNotEnoughTime() bool {
	return ((w >> 22) & 0x01) == 1
}

func (w *M9205Warn) SetRoadDriveNotEnoughTime(b bool) {
	if b {
		*w |= 1 << 22
	} else {
		*w &^= 1 << 22
	}
}

// RoadDeparture 路线偏离
func (w M9205Warn) RoadDeparture() bool {
	return ((w >> 23) & 0x01) == 1
}

func (w *M9205Warn) SetRoadDeparture(b bool) {
	if b {
		*w |= 1 << 23
	} else {
		*w &^= 1 << 23
	}
}

// VSSFault 车辆VSS故障
func (w M9205Warn) VSSFault() bool {
	return ((w >> 24) & 0x01) == 1
}

func (w *M9205Warn) SetVSSFault(b bool) {
	if b {
		*w |= 1 << 24
	} else {
		*w &^= 1 << 24
	}
}

// OilError 油量异常
func (w M9205Warn) OilError() bool {
	return ((w >> 25) & 0x01) == 1
}

func (w *M9205Warn) SetOilError(b bool) {
	if b {
		*w |= 1 << 25
	} else {
		*w &^= 1 << 25
	}
}

// VehStolen 车辆被盗
func (w M9205Warn) VehStolen() bool {
	return ((w >> 26) & 0x01) == 1
}

func (w *M9205Warn) SetVehStolen(b bool) {
	if b {
		*w |= 1 << 26
	} else {
		*w &^= 1 << 26
	}
}

// IllegalIgnition 非法点火
func (w M9205Warn) IllegalIgnition() bool {
	return ((w >> 27) & 0x01) == 1
}

func (w *M9205Warn) SetIllegalIgnition(b bool) {
	if b {
		*w |= 1 << 27
	} else {
		*w &^= 1 << 27
	}
}

// IllegalMove 非法位移
func (w M9205Warn) IllegalMove() bool {
	return ((w >> 28) & 0x01) == 1
}

func (w *M9205Warn) SetIllegalMove(b bool) {
	if b {
		*w |= 1 << 28
	} else {
		*w &^= 1 << 28
	}
}

// PreCrash 碰撞预警
func (w M9205Warn) PreCrash() bool {
	return ((w >> 29) & 0x01) == 1
}

func (w *M9205Warn) SetPreCrash(b bool) {
	if b {
		*w |= 1 << 29
	} else {
		*w &^= 1 << 29
	}
}

// PreRollOver 侧翻预警
func (w M9205Warn) PreRollOver() bool {
	return ((w >> 30) & 0x01) == 1
}

func (w *M9205Warn) SetPreRollOver(b bool) {
	if b {
		*w |= 1 << 30
	} else {
		*w &^= 1 << 30
	}
}

// IllegalOpenDoor 非法开门
func (w M9205Warn) IllegalOpenDoor() bool {
	return ((w >> 31) & 0x01) == 1
}

func (w *M9205Warn) SetIllegalOpenDoor(b bool) {
	if b {
		*w |= 1 << 31
	} else {
		*w &^= 1 << 31
	}
}

// VideoLoss 视频信号丢失报警
func (w M9205Warn) VideoLoss() bool {
	return ((w >> 32) & 0x01) == 1
}

func (w *M9205Warn) SetVideoLoss(value bool) {
	if value {
		*w |= 1 << 32
	} else {
		*w &^= 1 << 32
	}
}

// VideoShelter 视频信号遮挡报警
func (w M9205Warn) VideoShelter() bool {
	return ((w >> 33) & 0x01) == 1
}

func (w *M9205Warn) SetVideoShelter(value bool) {
	if value {
		*w |= 1 << 33
	} else {
		*w &^= 1 << 33
	}
}

// StorageFault 存储单元故障报警
func (w M9205Warn) StorageFault() bool {
	return ((w >> 34) & 0x01) == 1
}

func (w *M9205Warn) SetStorageFault(value bool) {
	if value {
		*w |= 1 << 34
	} else {
		*w &^= 1 << 34
	}
}

// OtherVideoFault 其他视频设备故障报警
func (w M9205Warn) OtherVideoFault() bool {
	return ((w >> 35) & 0x01) == 1
}

func (w *M9205Warn) SetOtherVideoFault(value bool) {
	if value {
		*w |= 1 << 35
	} else {
		*w &^= 1 << 35
	}
}

// BusOverload 客车超员报警
func (w M9205Warn) BusOverload() bool {
	return ((w >> 36) & 0x01) == 1
}

func (w *M9205Warn) SetBusOverload(value bool) {
	if value {
		*w |= 1 << 36
	} else {
		*w &^= 1 << 36
	}
}

// AbnormalDriving 异常驾驶行为报警
func (w M9205Warn) AbnormalDriving() bool {
	return ((w >> 37) & 0x01) == 1
}

func (w *M9205Warn) SetAbnormalDriving(value bool) {
	if value {
		*w |= 1 << 37
	} else {
		*w &^= 1 << 37
	}
}

// SpecialRecordFull 特殊报警录像达到存储阈值报警
func (w M9205Warn) SpecialRecordFull() bool {
	return ((w >> 38) & 0x01) == 1
}

func (w *M9205Warn) SetSpecialRecordFull(value bool) {
	if value {
		*w |= 1 << 38
	} else {
		*w &^= 1 << 38
	}
}
