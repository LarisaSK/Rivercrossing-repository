package state

func ViewState() string {
	return "[Kylling Rev Korn HS ---V \\_______/ _____________________Ø---]"
}
func PutInBoat(item string) string {
	return "[Rev Korn HS ---V \\___kylling___/_________________________Ø---]"
}

func CrossRiver() string {
	return "[Rev Korn HS ---V \\___________________________\\___kylling___/Ø---]"
}

//jeg skal ha noen if-statements på crossriver funksjon og put in boat funksjon som gjør at man
//eks ikke kan putte noe i båten hvis båten er i andre siden av elva eller hvis det er allerede fullt

//Du kan heller ikke cross river hvis du ikke har puttet noe i båten
//Det betyr at du må lage en variabel som hele tiden kjenner til tilstanden i spillet.
//Det betyr at den må oppdateres hver gang tilstanden endres. Dette må da PutInBoat funksjonen
//og CrossRiver registrere slik at feilmeldinger kommer opp når vi kjører disse funksjonene
// dersom tilstanden ikke passer
