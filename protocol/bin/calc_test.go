package bin

import (
	"testing"

	"github.com/mingkid/g-jtt/protocol/msg"
)

func TestCalculateMsgLength(t *testing.T) {
	m := new(msg.Msg[msg.M0100])
	m.Body.ProvincialID = 42
	m.Body.CityID = 2130
	m.Body.ManufacturerID = "12345"
	m.Body.TermModel = "00003624"
	m.Body.TermID = "0244487"
	m.Body.LicensePlateColor = 255
	m.Body.LicensePlate = "粤B88888"

	if size, _ := CalculateMsgLength(m.Body); size != 45 {
		t.Errorf("calculate msg length failed")
	}
}
