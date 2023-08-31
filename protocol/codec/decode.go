package codec

import (
	"errors"
	"reflect"
	"strconv"
	"strings"

	"github.com/mingkid/g-jtt/protocol/bin"
)

type Decoder struct {
	variable map[string]uint
}

func NewDecoder() *Decoder {
	return &Decoder{
		variable: make(map[string]uint),
	}
}

func (d *Decoder) Decode(msg interface{}, b []byte) error {
	msgType := reflect.TypeOf(msg)
	if msgType.Kind() != reflect.Ptr || msgType.Elem().Kind() != reflect.Struct {
		return errors.New("msg parameter must be a pointer to a struct")
	}

	reader := bin.NewReader(b)

	msgValue := reflect.ValueOf(msg).Elem()

	return d.decodeStruct(msgType.Elem(), &msgValue, reader)
}

func (d *Decoder) decodeStruct(msgType reflect.Type, msgValue *reflect.Value, reader *bin.Reader) error {
	// Iterate through the fields of the struct
	for i := 0; i < msgType.NumField(); i++ {
		fieldType := msgType.Field(i)
		fieldValue := msgValue.Field(i)

		tagValue := fieldType.Tag.Get("jtt13") // Change to the appropriate tag

		if tagValue == Ignore {
			continue
		}

		if fieldType.Type.Kind() == reflect.Struct {
			if err := d.decodeStruct(fieldType.Type, &fieldValue, reader); err != nil {
				return err
			}
			continue
		}

		fieldDecoder, err := NewFiledDecoder(fieldType, &fieldValue, tagValue, reader)
		if err != nil {
			return err
		}
		if err = fieldDecoder.Decode(); err != nil {
			return err
		}
	}

	return nil
}

func extractBCDLength(tagValue string) int {
	if parts := splitTag(tagValue); len(parts) == 2 && parts[0] == BCD {
		length, _ := strconv.Atoi(parts[1])
		return length
	}
	return 0
}

func splitTag(tagValue string) []string {
	return strings.Split(tagValue, ",")
}

const (
	Ignore = "-"
	BCD    = "bcd"
	Raw    = "raw"
	Var    = "var"
	UseVar = "$"
)
