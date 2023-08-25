package codec

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/mingkid/g-jtt/protocol/msg"
)

func TestDecodeM0200(t *testing.T) {
	dataHex := "0200006d036240156681017100000180000c0002020f9674069032f50168000000002308181859010104000076d930011b31011de60100e7080000000600000000f50101f6080861223060963705eb29000c00b2898604a21921c292847900060089ffffeffd000600c5feffffef0004002d0000000300a80e"
	data, _ := hex.DecodeString(dataHex)
	m := msg.New[msg.M0200]()

	var decoder Decoder
	_ = decoder.Decode(m, data)

	if m.Head.MsgID != msg.MsgIDTermLocationReport {
		t.Fatalf("消息包ID解析错误，应为%d，实际为%d", msg.MsgIDTermLocationReport, m.Head.MsgID)
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
	if m.Body.Status.GPSUsed() != true {
		t.Fatalf("消息包ID解析错误，应为%v，实际为%v", true, m.Body.Status.GPSUsed())
	}
}

func TestDecodeM0001(t *testing.T) {
	dataHex := "0001006d036240156681017101710200000000"
	data, _ := hex.DecodeString(dataHex)

	var decoder Decoder
	m := msg.New[msg.M0001]()

	_ = decoder.Decode(m, data)

	if m.Body.AnswerSerialNo != 369 {
		t.Fatalf("应答消息序号解析错误，应为%v，实际为%v", 369, m.Body.AnswerSerialNo)
	}
	if m.Body.Result != msg.M0001ResultOK {
		t.Fatalf("处理结果解析错误，应为%v，实际为%v", msg.M0001ResultOK, m.Body.Result)
	}
	if m.Body.AnswerMsgID != msg.MsgIDTermLocationReport {
		t.Fatalf("应答消息ID解析错误，应为%v，实际为%v", msg.MsgIDTermLocationReport, m.Body.AnswerMsgID)
	}
	if m.Body.ErrorCode != 0x0000 {
		t.Fatalf("错误代码解析错误，应为%v，实际为%v", 0x0000, m.Body.ErrorCode)
	}
}

func TestDecodeM0100(t *testing.T) {
	var decoder Decoder
	m := msg.New[msg.M0100]()

	_ = decoder.Decode(m, []byte{1, 0, 0, 39, 1, 48, 81, 25, 38, 117, 0, 128, 0, 44, 1, 44, 55, 48, 49, 49, 49, 66, 83, 74, 45, 65, 54, 66, 68, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 49, 49, 57, 50, 54, 55, 53, 0, 0, 0})

	if m.Body.ProvincialID != 44 {
		t.Fatalf("省份 ID 解析错误，应为%d，实际为%d", 44, m.Body.ProvincialID)
	}
	if m.Body.CityID != 300 {
		t.Fatalf("城市 ID 解析错误，应为%d，实际为%d", 300, m.Body.CityID)
	}
	if m.Body.ManufacturerID != "70111" {
		t.Fatalf("制造商 ID 解析错误，应为%s，实际为%s", "70111", m.Body.ManufacturerID)
	}
	if m.Body.TermModel != "BSJ-A6BD" {
		t.Fatalf("终端型号解析错误，应为%s，实际为%s", "BSJ-A6BD", m.Body.TermModel)
	}
	if m.Body.TermID != "1192675" {
		t.Fatalf("终端 ID 解析错误，应为%s，实际为%s", "1192675", m.Body.TermID)
	}
	if m.Body.LicensePlateColor != 0 {
		t.Fatalf("车牌颜色解析错误，应为%d，实际为%d", 0, m.Body.LicensePlateColor)
	}
	if m.Body.LicensePlate != "" {
		t.Fatalf("车牌号码解析错误，应为%s，实际为%s", "", m.Body.LicensePlate)
	}
}

func TestDecodeM0102(t *testing.T) {
	var (
		decoder Decoder
	)
	m := msg.New[msg.M0102]()

	_ = decoder.Decode(m, []byte{1, 0, 0, 39, 1, 48, 81, 25, 38, 117, 0, 128, 55, 54, 50, 54, 102, 53, 49, 57, 56, 48, 53, 51, 102, 98, 50, 99, 48, 49, 100, 100, 48, 101, 98, 101, 97, 100, 101, 54, 48, 99, 102, 51})

	if m.Body.Token != "7626f5198053fb2c01dd0ebeade60cf3" {
		t.Fatalf("Token 解析错误，应为%s，实际为%s", "7626f5198053fb2c01dd0ebeade60cf3", m.Body.Token)
	}
}

func TestDecodeM0704(t *testing.T) {
	var (
		m0200List []*msg.M0200
		d         Decoder
	)

	m := msg.New[msg.M0704]()

	//fmt.Println(hex.EncodeToString([]byte{7, 4, 0, 177, 1, 48, 80, 137, 135, 40, 4, 123, 0, 3, 1, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 125, 209, 177, 7, 86, 215, 109, 0, 0, 0, 0, 1, 24, 35, 8, 33, 18, 20, 1, 1, 4, 0, 6, 115, 122, 43, 4, 11, 51, 11, 51, 48, 1, 26, 49, 1, 25, 0, 4, 0, 206, 11, 51, 3, 2, 0, 0, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 125, 209, 175, 7, 86, 215, 132, 0, 0, 0, 33, 1, 24, 35, 8, 33, 18, 19, 49, 1, 4, 0, 6, 115, 122, 43, 4, 11, 41, 11, 41, 48, 1, 30, 49, 1, 27, 0, 4, 0, 206, 11, 41, 3, 2, 0, 33, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 125, 209, 174, 7, 86, 215, 114, 0, 0, 0, 0, 1, 24, 35, 8, 33, 18, 20, 49, 1, 4, 0, 6, 115, 122, 43, 4, 11, 51, 11, 51, 48, 1, 30, 49, 1, 27, 0, 4, 0, 206, 11, 51, 3, 2, 0, 0}))
	_ = d.Decode(m, []byte{7, 4, 0, 177, 1, 48, 80, 137, 135, 40, 4, 123, 0, 3, 1, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 125, 209, 177, 7, 86, 215, 109, 0, 0, 0, 0, 1, 24, 35, 8, 33, 18, 20, 1, 1, 4, 0, 6, 115, 122, 43, 4, 11, 51, 11, 51, 48, 1, 26, 49, 1, 25, 0, 4, 0, 206, 11, 51, 3, 2, 0, 0, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 125, 209, 175, 7, 86, 215, 132, 0, 0, 0, 33, 1, 24, 35, 8, 33, 18, 19, 49, 1, 4, 0, 6, 115, 122, 43, 4, 11, 41, 11, 41, 48, 1, 30, 49, 1, 27, 0, 4, 0, 206, 11, 41, 3, 2, 0, 33, 0, 56, 0, 0, 0, 0, 0, 12, 0, 3, 2, 125, 209, 174, 7, 86, 215, 114, 0, 0, 0, 0, 1, 24, 35, 8, 33, 18, 20, 49, 1, 4, 0, 6, 115, 122, 43, 4, 11, 51, 11, 51, 48, 1, 30, 49, 1, 27, 0, 4, 0, 206, 11, 51, 3, 2, 0, 0})

	items, _ := m.Body.Items()
	for _, item := range items {
		m0200 := new(msg.M0200)
		_ = d.Decode(m0200, item)
		m0200List = append(m0200List, m0200)
	}
	if m0200List[0].Latitude != 41800113 {
		t.Fatalf("纬度解析错误，应为%d，实际为%d", 41800113, m0200List[0].Latitude)
	}
	if m0200List[2].Latitude != 41800110 {
		t.Fatalf("纬度解析错误，应为%d，实际为%d", 41800113, m0200List[2].Latitude)
	}
}

func TestDecode(t *testing.T) {
	var (
		d Decoder
	)

	m := msg.New[msg.M8001]()

	b, _ := hex.DecodeString("8001000501305119491100000006010200")
	_ = d.Decode(m, b)
	fmt.Println("")
}
