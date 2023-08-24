package msg

import "testing"

func TestBodyProps(t *testing.T) {
	var b BodyProps = 0
	b.SetBodyLength(100)
	b.SetEncrypt()
	if b != 1124 {
		t.Errorf("设置消息体属性错误，应为 %d, 实际为 %d", 1124, b)
	}
}
