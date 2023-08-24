package bin

import (
	"testing"

	"github.com/mingkid/g-jtt/protocol/msg"
)

func TestCalculateMsgLength(t *testing.T) {
	m := msg.M0100{
		Head: msg.Head{},
		M0100Body: msg.M0100Body{
			ProvincialID:      42,
			CityID:            2130,
			ManufacturerID:    "12345",
			TermModel:         "00003624",
			TermID:            "0244487",
			LicensePlateColor: 255,
			LicensePlate:      "粤B88888",
		},
	}

	if size, _ := CalculateMsgLength(m.M0100Body); size != 45 {
		t.Errorf("calculate msg length failed")
	}
}
