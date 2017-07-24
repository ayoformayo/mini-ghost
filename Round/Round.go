package Round

// Round stuff
type Round struct {
	Number      int
	Fragment    string
	RoundStates []roundState
}

type roundState struct {
	Letter   string
	PlayerID int
}

// GameState does something
func (round *Round) GameState() string {
	compiledFragment := ""
	for _, roundState := range round.RoundStates {
		compiledFragment += roundState.Letter
	}

	return compiledFragment
}

// AppendLetter does something
func (round *Round) AppendLetter(letter string, PlayerID int) {
	round.RoundStates = append(round.RoundStates, roundState{Letter: letter, PlayerID: PlayerID})
}
