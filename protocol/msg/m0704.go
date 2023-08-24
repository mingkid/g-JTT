package msg

import (
	"encoding/binary"
	"fmt"
)

type M0704 struct {
	Head
	M0704Body
}

type M0704Body struct {
	Count    uint16 // 数据项个数
	Type     uint8  // 位置数据类型
	RawItems []byte // 位置汇报数据项，原始数据
	items    [][]byte
}

func (m *M0704Body) Items() (res [][]byte, err error) {
	if m.items == nil {
		if err = m.analyze(); err != nil {
			return
		}
	}
	return m.items, nil
}

func (m *M0704Body) analyze() error {
	if m.Count == 0 {
		return nil
	}

	rawItems := m.RawItems
	m.items = make([][]byte, m.Count)

	for i := uint16(0); i < m.Count; i++ {
		if len(rawItems) < 2 {
			return fmt.Errorf("not enough data for item length at index %d", i)
		}

		itemLen := binary.BigEndian.Uint16(rawItems[:2])
		rawItems = rawItems[2:]

		if len(rawItems) < int(itemLen) {
			return fmt.Errorf("not enough data for item content at index %d", i)
		}

		m.items[i] = rawItems[:itemLen]
		rawItems = rawItems[itemLen:]
	}

	return nil
}

type M0704Type uint8

const (
	// M0704TypeNormal 正常为止批量汇报
	M0704TypeNormal M0704Type = iota

	// M0704TypeAgain 盲区补报
	M0704TypeAgain
)
