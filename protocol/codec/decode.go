package codec

import "gihub.com/mingkid/g-jtt/protocol/msg"

type Decoder struct{}

func (d *Decoder) Decode(msg any, b []byte) error {
	return nil
}

func ExtraMsgID(b []byte) msg.MsgID {
	return 0
}
