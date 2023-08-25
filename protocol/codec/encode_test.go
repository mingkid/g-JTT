package codec

import (
	"encoding/hex"
	"fmt"
	"github.com/mingkid/g-jtt/protocol/msg"
	"testing"
)

func TestDecodeM8001(t *testing.T) {
	var e Encoder
	m := msg.M8001{
		Head: msg.Head{
			MsgID: msg.MsgIDPlatformCommResp,
			Phone: "13680179679",
			SN:    123,
		},
		M8001Body: msg.M8001{
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

func TestDecodeM8100(t *testing.T) {
	var e Encoder
	m := msg.M8100{
		Head: msg.Head{
			MsgID: msg.MsgIDTermRegResp,
			Phone: "13680179679",
			SN:    123,
		},
		M8100Body: msg.M8100{
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
