package bin

import (
	"encoding/hex"
	"strings"
)

const (
	IdentityBitChar       = "7e"   // 标识位
	EscapeBitChar         = "7d"   // 转义位
	IdentityBitEscapeChar = "7d02" // 7E 转义符
	EscapeBitEscapeChar   = "7d01" // 7D 转义符
)

// Unescape 返回反转义后的数据包，不影响原始数据包
func Unescape(b []byte) (res []byte) {
	str := hex.EncodeToString(b)

	// 去头尾
	str = strings.Trim(str, IdentityBitChar)

	// 反转义
	str = strings.Replace(str, IdentityBitEscapeChar, IdentityBitChar, -1)
	str = strings.Replace(str, EscapeBitEscapeChar, EscapeBitChar, -1)

	res, _ = hex.DecodeString(str)
	return
}

// Escape 返回
func Escape(b []byte) (res []byte) {
	str := hex.EncodeToString(b)

	// 转义
	str = strings.Replace(str, IdentityBitChar, IdentityBitEscapeChar, -1)
	str = strings.Replace(str, EscapeBitChar, EscapeBitEscapeChar, -1)

	// 加头尾
	str = IdentityBitChar + str + IdentityBitChar

	res, _ = hex.DecodeString(str)
	return
}
