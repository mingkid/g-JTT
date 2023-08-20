package codec

import (
	"encoding/hex"
	"testing"

	"gihub.com/mingkid/g-jtt/protocol/msg"
)

func TestDecodeM0200(t *testing.T) {
	dataHex := "0200006d036240156681017100000180000c0002020f9674069032f50168000000002308181859010104000076d930011b31011de60100e7080000000600000000f50101f6080861223060963705eb29000c00b2898604a21921c292847900060089ffffeffd000600c5feffffef0004002d0000000300a80e"
	data, _ := hex.DecodeString(dataHex)

	var (
		m       msg.M0200
		decoder Decoder
	)
	_ = decoder.Decode(&m, data)

	if m.Head.MsgID != msg.MsgIDTermLocationRepose {
		t.Fatalf("消息包ID解析错误，应为%d，实际为%d", msg.MsgIDTermLocationRepose, m.Head.MsgID)
	}
	if m.Head.BodyProps != 109 {
		t.Fatalf("消息包ID解析错误，应为%d，实际为%d", 109, m.Head.BodyProps)
	}
	if m.Head.Version != 0 {
		t.Fatalf("消息包ID解析错误，应为%d，实际为%d", 0, m.Head.Version)
	}
	if m.Head.Phone != "036240156681" {
		t.Fatalf("消息包ID解析错误，应为%s，实际为%s", "036240156681", m.Head.Phone)
	}
	if m.Head.SN != 369 {
		t.Fatalf("消息包ID解析错误，应为%d，实际为%d", 369, m.Head.SN)
	}
	//if m.Head.PackageInfo != nil {
	//	t.Fatalf("消息包ID解析错误，应为%v，实际为%v", nil, m.Head.PackageInfo)
	//}
	if m.Status.GPSUsed() != true {
		t.Fatalf("消息包ID解析错误，应为%v，实际为%v", true, m.Status.GPSUsed())
	}
}

func TestDecodeM0001(t *testing.T) {
	dataHex := "0001006d036240156681017101710200000000"
	data, _ := hex.DecodeString(dataHex)

	var (
		m       msg.M0001
		decoder Decoder
	)
	_ = decoder.Decode(&m, data)

	if m.AnswerSerialNo != 369 {
		t.Fatalf("应答消息序号解析错误，应为%v，实际为%v", 369, m.SN)
	}
	if m.Result != msg.M0001ResultOK {
		t.Fatalf("处理结果解析错误，应为%v，实际为%v", msg.M0001ResultOK, m.Result)
	}
	if m.AnswerMsgID != msg.MsgIDTermLocationRepose {
		t.Fatalf("应答消息ID解析错误，应为%v，实际为%v", msg.MsgIDTermLocationRepose, m.AnswerMsgID)
	}
	if m.ErrorCode != 0x0000 {
		t.Fatalf("错误代码解析错误，应为%v，实际为%v", 0x0000, m.ErrorCode)
	}
}
