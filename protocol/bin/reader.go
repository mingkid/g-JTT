package bin

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
)

type Reader struct {
	reader *bytes.Reader
}

func NewReader(data []byte) *Reader {
	return &Reader{
		reader: bytes.NewReader(data),
	}
}

func (r *Reader) ReadByte() (byte, error) {
	return r.reader.ReadByte()
}

func (r *Reader) Read(p []byte) (n int, err error) {
	return r.reader.Read(p)
}

func (r *Reader) ReadUint16() (uint16, error) {
	var data [2]byte
	_, err := r.reader.Read(data[:])
	if err != nil {
		return 0, err
	}
	return uint16(data[0])<<8 | uint16(data[1]), nil
}

func (r *Reader) ReadBCD(length int) (string, error) {
	data := make([]byte, length)
	_, err := r.reader.Read(data)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", data), nil
}

func (r *Reader) ReadUint32() (uint32, error) {
	var data [4]byte
	_, err := r.reader.Read(data[:])
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(data[:]), nil
}

func (r *Reader) ReadBytes(size int) ([]byte, error) {
	data := make([]byte, size)
	_, err := r.reader.Read(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Reader) ReadB() (byte, error) {
	return r.reader.ReadByte()
}

func (r *Reader) ReadString(size int) (string, error) {
	data := make([]byte, size)
	_, err := r.reader.Read(data)
	if err != nil {
		return "", err
	}

	// Convert GBK encoded bytes to string
	decoder := simplifiedchinese.GBK.NewDecoder()
	decodedData, _, err := transform.Bytes(decoder, data)
	if err != nil {
		return "", err
	}

	return string(decodedData), nil
}

func (r *Reader) ReadStringAll() (string, error) {
	remaining := r.Remaining()
	return r.ReadString(remaining)
}

func (r *Reader) ReadAll() ([]byte, error) {
	return io.ReadAll(r.reader)
}

// Remaining returns the number of remaining bytes to read
func (r *Reader) Remaining() int {
	return int(r.reader.Size()) - int(r.reader.Len())
}
