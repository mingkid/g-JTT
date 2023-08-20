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

func TestDecodeM0100(t *testing.T) {
	var (
		m       msg.M0100
		decoder Decoder
	)
	_ = decoder.Decode(&m, []byte{1, 0, 0, 39, 1, 48, 81, 25, 38, 117, 0, 128, 0, 44, 1, 44, 55, 48, 49, 49, 49, 66, 83, 74, 45, 65, 54, 66, 68, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 49, 49, 57, 50, 54, 55, 53, 0, 0, 0})

	if m.ProvincialID != 44 {
		t.Fatalf("省份 ID 解析错误，应为%d，实际为%d", 44, m.ProvincialID)
	}
	if m.CityID != 300 {
		t.Fatalf("城市 ID 解析错误，应为%d，实际为%d", 300, m.CityID)
	}
	if m.ManufacturerID != "70111" {
		t.Fatalf("制造商 ID 解析错误，应为%s，实际为%s", "70111", m.ManufacturerID)
	}
	if m.TermModel != "BSJ-A6BD" {
		t.Fatalf("终端型号解析错误，应为%s，实际为%s", "BSJ-A6BD", m.TermModel)
	}
	if m.TermID != "1192675" {
		t.Fatalf("终端 ID 解析错误，应为%s，实际为%s", "1192675", m.TermID)
	}
	if m.LicensePlateColor != 0 {
		t.Fatalf("车牌颜色解析错误，应为%d，实际为%d", 0, m.LicensePlateColor)
	}
	if m.LicensePlate != "" {
		t.Fatalf("车牌号码解析错误，应为%s，实际为%s", "", m.LicensePlate)
	}
}

func TestDecodeM0102(t *testing.T) {
	var (
		m       msg.M0102
		decoder Decoder
	)
	_ = decoder.Decode(&m, []byte{1, 0, 0, 39, 1, 48, 81, 25, 38, 117, 0, 128, 55, 54, 50, 54, 102, 53, 49, 57, 56, 48, 53, 51, 102, 98, 50, 99, 48, 49, 100, 100, 48, 101, 98, 101, 97, 100, 101, 54, 48, 99, 102, 51})

	if m.Token != "7626f5198053fb2c01dd0ebeade60cf3" {
		t.Fatalf("Token 解析错误，应为%s，实际为%s", "7626f5198053fb2c01dd0ebeade60cf3", m.Token)
	}
}
