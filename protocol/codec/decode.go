package codec

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"gihub.com/mingkid/g-jtt/protocol/bin"
)

type Decoder struct{}

func (d *Decoder) Decode(msg interface{}, b []byte) error {
	msgType := reflect.TypeOf(msg)
	if msgType.Kind() != reflect.Ptr || msgType.Elem().Kind() != reflect.Struct {
		return errors.New("msg parameter must be a pointer to a struct")
	}

	reader := bin.NewReader(b)

	msgValue := reflect.ValueOf(msg).Elem()

	// Iterate through the fields of the struct
	for i := 0; i < msgType.Elem().NumField(); i++ {
		fieldType := msgType.Elem().Field(i)
		fieldValue := msgValue.Field(i)

		tagValue := fieldType.Tag.Get("jtt13") // Change to the appropriate tag

		if tagValue == "-" {
			continue
		}

		switch fieldType.Type.Kind() {
		case reflect.Struct:
			if err := d.Decode(fieldValue.Addr().Interface(), b); err != nil {
				return err
			}

		case reflect.Uint8:
			val, err := reader.ReadByte()
			if err != nil {
				return err
			}
			fieldValue.SetUint(uint64(val))

		case reflect.Uint16:
			val, err := reader.ReadUint16()
			if err != nil {
				return err
			}
			fieldValue.SetUint(uint64(val))

		case reflect.Uint32:
			val, err := reader.ReadUint32()
			if err != nil {
				return err
			}
			fieldValue.SetUint(uint64(val))

		case reflect.String:
			if tagValue == "" {
				// Read remaining data and convert to string
				val, err := reader.ReadStringAll()
				if err != nil {
					return err
				}
				fieldValue.SetString(val)
			} else if tagValue != "-" {
				if bcdLength := extractBCDLength(tagValue); bcdLength > 0 {
					val, err := reader.ReadBCD(bcdLength)
					if err != nil {
						return err
					}
					fieldValue.SetString(val)
				} else {
					// Read specified size of data
					size, _ := strconv.Atoi(tagValue)
					val, err := reader.ReadString(size)
					if err != nil {
						return err
					}
					fieldValue.SetString(val)
				}
			}

		// Add cases for other supported types

		default:
			// Unsupported field type
			return fmt.Errorf("unsupported field type: %v", fieldType.Type.Kind())
		}
	}

	return nil
}

func extractBCDLength(tagValue string) int {
	if parts := splitTag(tagValue); len(parts) == 2 && parts[0] == "bcd" {
		length, _ := strconv.Atoi(parts[1])
		return length
	}
	return 0
}

func splitTag(tagValue string) []string {
	return strings.Split(tagValue, ",")
}
