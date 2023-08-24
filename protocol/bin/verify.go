package bin

import "fmt"

// Checksum 返回校验和
func Checksum(b []byte) (sum byte) {
	sum = b[0] // 取第0个，从第1个开始异或
	for i := 1; i < len(b); i++ {
		sum = sum ^ b[i]
	}
	return
}

// Verify 检验校验码
func Verify(b []byte, c uint8) error {
	if Checksum(b) != c {
		return fmt.Errorf("checksum error")
	}
	return nil
}
