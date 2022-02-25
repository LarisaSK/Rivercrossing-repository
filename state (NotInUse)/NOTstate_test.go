package state

import (
	"testing"
)

func TestViewState(t *testing.T) {
	wanted := "[Kylling Rev Korn HS ---V \\_______/ _____________________Ø---]"
	got := ViewState()
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}
func TestPutInBoat(t *testing.T) {
	wanted := "[Rev Korn HS ---V \\___kylling___/_________________________Ø---]"
	got := PutInBoat("Kylling")
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}
func TestCrossRiver(t *testing.T) {
	wanted := "[Rev Korn HS ---V \\___________________________\\___kylling___/Ø---]"
	got := CrossRiver()
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}
