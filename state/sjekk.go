package state

import (
	"fmt"
)

func main() {

	//////////////////////////////////////////////////

	var currentState int
	var states [15]string
	states[0] = "HS KY KO RE V--\\_______/_______/--Ø"
	states[1] = "HS    KO RE V--\\__KY___/_______/--Ø"
	states[2] = "      KO RE V--\\_HS_KY_/_______/--Ø"
	states[3] = "      KO RE V--\\________\\HS_KY_/--Ø"
	states[4] = "      KO RE V--\\__HS___/_______/--Ø KY"
	states[5] = "      KO    V--\\__HS_RE_/_______/--Ø KY"
	states[6] = "      KO    V--\\________\\HS_RE_/--Ø KY"
	states[7] = "      KO    V--\\________\\HS_KY_/--Ø RE"
	states[8] = "      KO    V--\\__HS_KY_/_______/--Ø RE"
	states[9] = "   KY       V--\\__HS_KO_/_______/--Ø RE"
	states[10] = "   KY       V--\\________\\HS_KO_/--Ø RE"
	states[11] = "   KY       V--\\__HS___/_______/--Ø RE KO"
	states[12] = "            V--\\__HS_KY_/_______/--Ø RE KO"
	states[13] = "            V--\\_______\\__HS_KY_/--Ø RE KO"
	states[14] = "            V--\\_______\\_______/--Ø RE KO HS KY"

	/////////////////////////////////////////////////////

	currentState = 0

	ViewCurrentState(states, currentState)
	currentState = PutInBoat(currentState, "KY")
	ViewCurrentState(states, currentState)
	currentState = GetHSInBoat(currentState)
	ViewCurrentState(states, currentState)
	currentState = CrossRiver(currentState, "V--\\_HS_KY_/")
	ViewCurrentState(states, currentState)
	currentState = CrossRiver(currentState, "\\HS_KY_/--Ø")
	ViewCurrentState(states, currentState)
	currentState = PutInBoat(currentState, "RE")
	ViewCurrentState(states, currentState)
	currentState = CrossRiver(currentState, "V--\\__HS_RE_/")
	ViewCurrentState(states, currentState)
	currentState = PutInBoat(currentState, "KY")
	ViewCurrentState(states, currentState)
	currentState = CrossRiver(currentState, "\\HS_KY_/--Ø RE")
	ViewCurrentState(states, currentState)
	currentState = PutInBoat(currentState, "KO")
	ViewCurrentState(states, currentState)
	currentState = CrossRiver(currentState, "V--\\__HS_KO_/")
	ViewCurrentState(states, currentState)
	currentState = CrossRiver(currentState, "\\HS_KO_/--Ø RE")
	ViewCurrentState(states, currentState)
	currentState = PutInBoat(currentState, "KY")
	ViewCurrentState(states, currentState)
	currentState = CrossRiver(currentState, "V--\\__HS_KY_/")
	ViewCurrentState(states, currentState)
	currentState = GetOutOfBoat(currentState, "\\__HS_KY_/--Ø RE KO")
	ViewCurrentState(states, currentState)

}

//////////////////////////////////////////////////

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
