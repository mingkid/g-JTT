package codec

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"

	"github.com/mingkid/g-jtt/protocol/bin"
)

// FieldDecoder 字段解码器
type FieldDecoder interface {
	Decode() error
}

// NumericFieldDecoder  数字字段解码器
type NumericFieldDecoder struct {
	fv       *reflect.Value
	readFunc func() (uint64, error)
}

func (d *NumericFieldDecoder) Decode() error {
	val, err := d.readFunc()
	if err != nil {
		return err
	}
	d.fv.SetUint(val)
	return nil
}

// SliceFieldDecoder 切片字段解码器
type SliceFieldDecoder struct {
	fv     *reflect.Value
	r      *bin.Reader
	tagVal string
}

func (d *SliceFieldDecoder) Decode() error {
	var (
		val []byte
		err error
	)

	if d.tagVal == "" {
		// Read remaining data
		val, err = d.r.ReadAll()
		if err != nil {
			return err
		}
	} else {
		decodeType, length, err := extractLength(d.tagVal)
		if err != nil {
			return err
		}
		if decodeType == CodecTypeRaw {
			val, err = d.r.ReadBytes(length)
			if err != nil {
				return err
			}
		}
	}
	d.fv.SetBytes(val)
	return nil
}

// MapFieldDecoder  字典字段解码器
type MapFieldDecoder struct {
	f  reflect.StructField
	fv *reflect.Value
	r  *bin.Reader
}

func (d *MapFieldDecoder) Decode() error {
	mapValue := reflect.MakeMap(d.f.Type)
	for {
		key, err := d.r.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		size, err := d.r.ReadByte()
		if err != nil {
			return err
		}

		val, err := d.r.ReadBytes(int(size))
		if err != nil {
			return err
		}

		mapValue.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(val))
	}
	d.fv.Set(mapValue)
	return nil
}

// StringFieldDecoder  字符串字段解码器
type StringFieldDecoder struct {
	fv      *reflect.Value
	r       *bin.Reader
	tagVal  string
	globVar map[string]uint
}

func (d *StringFieldDecoder) Decode() error {
	if d.tagVal == "" {
		return d.remainingDecode()
	}
	decodeType, length, err := extractLength(d.tagVal)
	if err == nil {
		if decodeType == CodecTypeBCD {
			return d.bcdDecode(length)
		}
	}

	return d.sizeDecode()
}

// sizeDecode 解码指定长度的字符串
func (d *StringFieldDecoder) sizeDecode() error {
	size, _ := strconv.Atoi(d.tagVal)
	val, err := d.r.ReadString(size)
	if err != nil {
		return err
	}
	d.fv.SetString(val)
	return nil
}

// bcdDecode 解码BCD编码的字符串
func (d *StringFieldDecoder) bcdDecode(bcdLength int) error {
	val, err := d.r.ReadBCD(bcdLength)
	if err != nil {
		return err
	}
	d.fv.SetString(val)
	return nil
}

// remainingDecode 读取剩余数据
func (d *StringFieldDecoder) remainingDecode() error {
	val, err := d.r.ReadStringAll()
	if err != nil {
		return err
	}
	d.fv.SetString(val)
	return nil
}

func NewFiledDecoder(f reflect.StructField, fv *reflect.Value, tagVal string, r *bin.Reader, globVar map[string]uint) (FieldDecoder, error) {
	switch fv.Kind() {
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		d := &NumericFieldDecoder{fv: fv}
		switch fv.Kind() {
		case reflect.Uint8:
			d.readFunc = func() (uint64, error) {
				val, err := r.ReadByte()
				return uint64(val), err
			}
		case reflect.Uint16:
			d.readFunc = func() (uint64, error) {
				val, err := r.ReadUint16()
				return uint64(val), err
			}
		case reflect.Uint32:
			d.readFunc = func() (uint64, error) {
				val, err := r.ReadUint32()
				return uint64(val), err
			}
		}
		return d, nil

	case reflect.Slice:
		if f.Type.Elem().Kind() != reflect.Uint8 {
			return nil, newFieldDecodeError(f.Name, fmt.Sprintf("不支持的数据类型 %s", f.Type.Elem().Kind().String()))
		}
		return &SliceFieldDecoder{
			fv:     fv,
			r:      r,
			tagVal: tagVal,
		}, nil

	case reflect.Map:
		if kind := f.Type.Elem().Kind(); kind != reflect.Slice {
			return nil, fmt.Errorf("不支持的数据类型 %s", kind)
		}
		if kind := f.Type.Elem().Elem().Kind(); kind != reflect.Uint8 {
			return nil, fmt.Errorf("不支持的数据类型 %s", kind)
		}
		return &MapFieldDecoder{
			f:  f,
			fv: fv,
			r:  r,
		}, nil

	case reflect.String:
		return &StringFieldDecoder{
			fv:      fv,
			r:       r,
			tagVal:  tagVal,
			globVar: globVar,
		}, nil

	default:
		return nil, fmt.Errorf("不支持的数据类型 %s", f.Type.Kind())
	}
}

func extractLength(tagValue string) (CodecType, int, error) {
	parts := strings.Split(tagValue, ",")
	if len(parts) != 2 {
		return "", 0, errors.New("没有可提取的长度信息")
	}
	length, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, errors.New("没有可提取的长度信息")
	}
	return CodecType(parts[0]), length, err
}

// fieldDecodeError 字段解码异常
type fieldDecodeError struct {
	s         string
	fieldName string
}

func newFieldDecodeError(filedName, text string) fieldDecodeError {
	return fieldDecodeError{
		s:         text,
		fieldName: filedName,
	}
}

func (err fieldDecodeError) Error() string {
	return fmt.Sprintf("%s 字段解码异常: %s", err.fieldName, err.s)
}
