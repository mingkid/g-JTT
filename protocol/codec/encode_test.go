package codec

import (
	"encoding/hex"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/mingkid/g-jtt/protocol/msg"
)

func TestEncodeM8001(t *testing.T) {
	var e Encoder
	m := msg.Msg[msg.M8001]{
		Head: msg.Head{
			MsgID: msg.MsgIDPlatformCommResp,
			Phone: "13680179679",
			SN:    123,
		},
		Body: msg.M8001{
			AnswerSerialNo: 321,
			AnswerMsgID:    msg.MsgIDTermLocationReport,
			Result:         msg.M8001ResultFail,
		},
	}

	b, _ := e.Encode(m)
	fmt.Println(hex.EncodeToString(b))
	if hex.EncodeToString(b) != "80010000013680179679007b0141020001" {
		t.Fatalf("组包错误，应为%s，实际为%s", "80010000013680179679007b0141020001", hex.EncodeToString(b))
	}
}

func TestEncodeM8100(t *testing.T) {
	var e Encoder
	m := msg.Msg[msg.M8100]{
		Head: msg.Head{
			MsgID: msg.MsgIDTermRegResp,
			Phone: "13680179679",
			SN:    123,
		},
		Body: msg.M8100{
			AnswerSerialNo: 321,
			Result:         msg.M8100ResultSuccess,
			Token:          "1234567890",
		},
	}

	b, _ := e.Encode(m)
	if hex.EncodeToString(b) != "81000000013680179679007b01410031323334353637383930" {
		t.Fatalf("组包错误，应为%s，实际为%s", "81000000013680179679007b01410031323334353637383930", hex.EncodeToString(b))
	}
}

func TestEncodeM9101(t *testing.T) {
	var e Encoder
	m := msg.M9101{
		LogicalChannel: 3,
		DataType:       msg.M9101DataTypeAudioVideo,
		StreamType:     msg.M9101StreamTypeMain,
	}
	m.SetTCPAddr(net.ParseIP("192.168.0.123"), 123)
	m.SetUDPAddr(net.ParseIP("192.168.0.123"), 456)

	b, _ := e.Encode(m)

	if hex.EncodeToString(b) != "0d3139322e3136382e302e313233007b01c8030000" {
		t.Fatalf("组包错误，应为%s，实际为%s", "0d3139322e3136382e302e313233007b01c8030000", hex.EncodeToString(b))
	}
}

func TestEncodeM9102(t *testing.T) {
	var e Encoder
	m := msg.M9102{
		LogicChannelNumber:  1,
		ControlDirective:    msg.M9102ControlSwitchStream,
		CloseAudioVideoType: msg.M9102CloseAudio,
		SwitchStreamType:    msg.M9102SwitchToSubStream,
	}

	b, _ := e.Encode(m)

	if hex.EncodeToString(b) != "01010101" {
		t.Fatalf("组包错误，应为%s，实际为%s", "01010101", hex.EncodeToString(b))
	}
}

func TestEncodeM9205(t *testing.T) {
	var e Encoder
	m := msg.NewM9205(4)
	startTime, _ := time.Parse("20060102150405", "20230904000000")
	m.SetStartTime(startTime)
	endTime, _ := time.Parse("20060102150405", "20230904171002")
	m.SetEndTime(endTime)
	m.Warn.SetOverSpeed(true)
	m.AVResourceType = msg.M9205AVResourceVideo
	m.StreamType = msg.M9205StreamMain
	m.StorageType = msg.M9205StorageTypeMain

	b, _ := e.Encode(m)

	if hex.EncodeToString(b) != "042309040000002309041710020000000000000002020101" {
		t.Fatalf("组包错误，应为%s，实际为%s", "042309040000002309041710020000000000000002020101", hex.EncodeToString(b))
	}
}
