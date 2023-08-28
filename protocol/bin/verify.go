package bin

import "fmt"

// Checksum 返回校验和
func Checksum(b []byte) (sum byte) {
	for _, i := range b {
		sum ^= i
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
