package msg

type M9105 struct {
	LogicChannelNumber uint8 // 逻辑通道号
	PacketLoss         uint8 // 丢包率，数值乘以 100 后取整
}
