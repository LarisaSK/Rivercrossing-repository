package main

import (
	"github.com/LarisaSK/Rivercrossing-repository/state"
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

	state.ViewCurrentState(states, currentState)
	currentState = state.PutInBoat(currentState, "KY")
	state.ViewCurrentState(states, currentState)
	currentState = state.GetHSInBoat(currentState)
	state.ViewCurrentState(states, currentState)
	currentState = state.CrossRiver(currentState, "V--\\_HS_KY_/")
	state.ViewCurrentState(states, currentState)
	currentState = state.CrossRiver(currentState, "\\HS_KY_/--Ø")
	state.ViewCurrentState(states, currentState)
	currentState = state.PutInBoat(currentState, "RE")
	state.ViewCurrentState(states, currentState)
	currentState = state.CrossRiver(currentState, "V--\\__HS_RE_/")
	state.ViewCurrentState(states, currentState)
	currentState = state.PutInBoat(currentState, "KY")
	state.ViewCurrentState(states, currentState)
	currentState = state.CrossRiver(currentState, "\\HS_KY_/--Ø RE")
	state.ViewCurrentState(states, currentState)
	currentState = state.PutInBoat(currentState, "KO")
	state.ViewCurrentState(states, currentState)
	currentState = state.CrossRiver(currentState, "V--\\__HS_KO_/")
	state.ViewCurrentState(states, currentState)
	currentState = state.CrossRiver(currentState, "\\HS_KO_/--Ø RE")
	state.ViewCurrentState(states, currentState)
	currentState = state.PutInBoat(currentState, "KY")
	state.ViewCurrentState(states, currentState)
	currentState = state.CrossRiver(currentState, "V--\\__HS_KY_/")
	state.ViewCurrentState(states, currentState)
	currentState = state.GetOutOfBoat(currentState, "\\__HS_KY_/--Ø RE KO")
	state.ViewCurrentState(states, currentState)

}
