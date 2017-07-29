package Round

import "github.com/ayoformayo/mini-ghost/Dictionary"

// Round stuff
type Round struct {
	Number      int
	Fragment    string
	RoundStates []roundState
	Dictionary  *Dictionary.Dictionary
}

type roundState struct {
	Letter   string
	PlayerID int
}

// IsOver determines if round has ended
func (round *Round) IsOver() bool {
	return round.Dictionary.WordTree.FragmentIsWord(round.GameState())
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
