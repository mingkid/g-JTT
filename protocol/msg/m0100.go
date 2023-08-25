package msg

type M0100 struct {
	ProvincialID      uint16 // 省份 ID
	CityID            uint16 // 城市 ID
	ManufacturerID    string `jtt13:"5"`  // 制造商 ID
	TermModel         string `jtt13:"20"` // 终端型号
	TermID            string `jtt13:"7"`  // 终端序列号
	LicensePlateColor byte   // 车牌号码颜色
	LicensePlate      string // 车牌号码
}
