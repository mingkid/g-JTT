package bin

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func CalculateMsgLength(msg any) (size uint16, err error) {
	// 遍历结构体的字段，计算字段的大小
	messageValue := reflect.ValueOf(msg)
	for i := 0; i < messageValue.NumField(); i++ {
		field := messageValue.Field(i)
		fieldType := messageValue.Type().Field(i)

		// 从标签中获取大小信息
		sizeTag := fieldType.Tag.Get("jtt13")
		if sizeTag == "-" {
			continue
		}

		// 根据类型和标签规则计算大小
		fieldSize, err := calculateFieldSize(field, sizeTag)
		if err != nil {
			return 0, fmt.Errorf("type %s ", err.Error())
		}

		size += uint16(fieldSize)
	}
	return
}

// calculateFieldSize 根据字段类型和标签规则计算字段大小
func calculateFieldSize(fieldValue reflect.Value, jttTag string) (int, error) {
	switch fieldValue.Kind() {
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return binary.Size(fieldValue.Interface()), nil

	case reflect.String:
		// BCD 编码
		if strings.HasPrefix(jttTag, "bcd") {
			parts := strings.Split(jttTag, ",")
			return strconv.Atoi(parts[1])
		}

		// GBK 编码
		size := 0
		if _, err := fmt.Sscanf(jttTag, "%d", &size); err == nil {
			// 有长度限制，例如 jtt13:"10"
			return size, nil
		} else {
			//  无长度限制，即无 jtt 标签
			if gbkBytes, _, err := transform.Bytes(simplifiedchinese.GBK.NewEncoder(), []byte(fieldValue.String())); err != nil {
				return 0, err
			} else {
				return len(gbkBytes), nil
			}
		}

	case reflect.Map:
		if fieldValue.Type().Elem().Kind() == reflect.Slice && fieldValue.Type().Elem().Elem().Kind() == reflect.Uint8 {
			size := 0
			mapKeys := fieldValue.MapKeys()
			for _, key := range mapKeys {
				// 如 M0200Extract ,附加消息 ID（1 位）+ 附加消息长度（1 位）+ 附加消息（n=附加消息长度）
				size += 2 + len(fieldValue.MapIndex(key).Bytes())
			}
		}
		return 0, fmt.Errorf("field types(map[%s][]%s) not supported for calculating field lengths",
			fieldValue.Type().String(),
			fieldValue.Type().Elem().String(),
		)

	// 根据实际情况添加其他类型的处理逻辑
	default:
		return 0, fmt.Errorf("field types(%s) not supported for calculating field lengths",
			fieldValue.Type().String(),
		)
	}
}
