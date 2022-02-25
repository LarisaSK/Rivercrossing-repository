package state

import (
	"fmt"
)

func GetHSInBoat(currentState int) int {
	if currentState == 1 {
		currentState = 2
	} else if currentState == 0 {
		currentState = -1
	}
	return currentState
}

/////////////////////////////////////////////////

func PutInBoat(currentState int, item string) int {
	if currentState == 0 {
		if item == "KY" {
			currentState = 1
		}
	}
	if currentState == 4 {
		if item == "RE" {
			currentState = 5
		}
	}
	if currentState == 6 {
		if item == "KY" {
			currentState = 7
		}

	}
	if currentState == 8 {
		if item == "KO" {
			currentState = 9
		}
	}
	if currentState == 11 {
		if item == "KY" {
			currentState = 12
		}
	}
	return currentState
}

///////////////////////////////////////////////////

func CrossRiver(currentState int, item string) int {
	if currentState == 2 {
		if item == "V--\\_HS_KY_/" {
			currentState = 3
		}

	}
	if currentState == 3 {
		if item == "\\HS_KY_/--Ø" {
			currentState = 4
		}
	}
	if currentState == 5 {
		if item == "V--\\__HS_RE_/" {
			currentState = 6
		}
	}
	if currentState == 7 {
		if item == "\\HS_KY_/--Ø RE" {
			currentState = 8
		}
	}
	if currentState == 9 {
		if item == "V--\\__HS_KO_/" {
			currentState = 10
		}
	}
	if currentState == 10 {
		if item == "\\HS_KO_/--Ø RE" {
			currentState = 11
		}
	}
	if currentState == 12 {
		if item == "V--\\__HS_KY_/" {
			currentState = 13
		}
	}
	return currentState
}

////////////////////////////////////////////////////

func GetOutOfBoat(currentState int, item string) int {
	if currentState == 13 {
		if item == "\\__HS_KY_/--Ø RE KO" {
			currentState = 14
		}
	}
	return currentState
}

////////////////////////////////////////////////////////
func ViewCurrentState(allStates [15]string, currentState int) {
	fmt.Println(allStates[currentState])
}
