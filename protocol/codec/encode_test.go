package codec

import (
	"encoding/hex"
	"fmt"
	"gihub.com/mingkid/g-jtt/protocol/msg"
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
		SerialNumber: 321,
		MsgID:        msg.MsgIDTermLocationRepose,
		Result:       0,
	}

	b, _ := e.Encode(m)
	fmt.Println(hex.EncodeToString(b))
}
