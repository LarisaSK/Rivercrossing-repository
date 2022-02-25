package state

import (
	"testing"
)

func TestGetHSInBoat(t *testing.T) {
	got := GetHSInBoat(2)
	t.Log(got)
}
func TestPutInBoat(t *testing.T) {
	got := PutInBoat(1, "KY")
	t.Log(got)

}

func TestCrossRiver(t *testing.T) {
	got := CrossRiver(3, "\\HS_KY_/--Ø")
	t.Log(got)

}

func TestGetOutOfBoat(t *testing.T) {
	got := GetOutOfBoat(14, "--Ø RE KO HS KY")
	t.Log(got)
}
