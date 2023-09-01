package codec

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/mingkid/g-jtt/protocol/bin"
)

type Encoder struct{}

func (e *Encoder) Encode(msg any) ([]byte, error) {
	writer := bin.NewWriter()

	msgType := reflect.ValueOf(msg)
	if msgType.Kind() == reflect.Ptr {
		msgType = msgType.Elem()
	}

	err := e.encodeStruct(msgType, writer)
	if err != nil {
		return nil, err
	}

	return writer.Bytes(), nil
}

func (e *Encoder) encodeStruct(structValue reflect.Value, writer *bin.Writer) error {
	structType := structValue.Type()

	for i := 0; i < structType.NumField(); i++ {
		fieldValue := structValue.Field(i)
		fieldType := structType.Field(i)

		tagValue := fieldType.Tag.Get("jtt13")

		if tagValue == Ignore {
			continue
		}

		switch fieldValue.Type().Kind() {
		case reflect.Uint8:
			if err := writer.WriteUint8(uint8(fieldValue.Uint())); err != nil {
				return err
			}

		case reflect.Uint16:
			if err := writer.WriteUint16(uint16(fieldValue.Uint())); err != nil {
				return err
			}

		case reflect.String:
			if strings.HasPrefix(tagValue, BCD) {
				// BCD 编码
				parts := strings.Split(tagValue, ",")
				length, _ := strconv.Atoi(parts[1])
				err := writer.WriteBCD(fieldValue.String(), length)
				if err != nil {
					return err
				}

			} else {
				//  GBK 编码
				err := writer.WriteString(fieldValue.String())
				if err != nil {
					return err
				}
			}

		case reflect.Map:
			if fieldValue.Type().Elem().Kind() == reflect.Slice && fieldValue.Type().Elem().Elem().Kind() == reflect.Uint8 {
				mapKeys := fieldValue.MapKeys()
				for _, key := range mapKeys {
					keyVal := key.Uint()
					valVal := fieldValue.MapIndex(key).Bytes()

					writer.WriteByte(byte(keyVal))
					writer.WriteByte(byte(len(valVal)))
					writer.Write(valVal)
				}
			}

		case reflect.Struct:
			err := e.encodeStruct(fieldValue, writer)
			if err != nil {
				return err
			}

		default:
			return fmt.Errorf("unsupported field type: %v", fieldValue.Type().Kind())
		}
	}

	return nil
}
