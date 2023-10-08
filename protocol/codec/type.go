package codec

// CodecType 编解码方式
type CodecType string

func (t CodecType) String() string {
	return string(t)
}

const (
	CodecTypeBCD CodecType = "bcd"
	CodecTypeRaw CodecType = "raw"
)
