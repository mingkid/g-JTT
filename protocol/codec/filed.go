package codec

import (
	"fmt"
	"io"
	"reflect"
	"strconv"

	"github.com/mingkid/g-jtt/protocol/bin"
)

type FieldDecoder interface {
	Decode() error
}

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
		if parts := splitTag(d.tagVal); len(parts) == 2 && parts[0] == Raw {
			size, _ := strconv.Atoi(parts[1])
			val, err = d.r.ReadBytes(size)
			if err != nil {
				return err
			}
			d.fv.SetBytes(val)
		}
	}
	d.fv.SetBytes(val)
	return nil
}

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

type StringFieldDecoder struct {
	fv     *reflect.Value
	r      *bin.Reader
	tagVal string
}

func (d *StringFieldDecoder) Decode() error {
	if d.tagVal == "" {
		// Read remaining data and convert to string
		val, err := d.r.ReadStringAll()
		if err != nil {
			return err
		}
		d.fv.SetString(val)
	} else {
		if bcdLength := extractBCDLength(d.tagVal); bcdLength > 0 {
			val, err := d.r.ReadBCD(bcdLength)
			if err != nil {
				return err
			}
			d.fv.SetString(val)
		} else {
			// Read specified size of data
			size, _ := strconv.Atoi(d.tagVal)
			val, err := d.r.ReadString(size)
			if err != nil {
				return err
			}
			d.fv.SetString(val)
		}
	}
	return nil
}

func NewFiledDecoder(f reflect.StructField, fv *reflect.Value, tagVal string, r *bin.Reader) (FieldDecoder, error) {
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
			return nil, FiledCantDecodeError{s: fmt.Sprintf("不支持的数据类型 %s []%s", f.Name, f.Type.Elem().Kind().String())}
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
			fv:     fv,
			r:      r,
			tagVal: tagVal,
		}, nil

	default:
		return nil, fmt.Errorf("不支持的数据类型 %s", f.Type.Kind())
	}
}

// FiledCantDecodeError 字段解码异常
type FiledCantDecodeError struct {
	s string
}

func (e FiledCantDecodeError) Error() string {
	return e.s
}
