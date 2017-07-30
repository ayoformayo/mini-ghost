package Round

import (
	"github.com/ayoformayo/mini-ghost/Dictionary"
)

// Round stuff
type Round struct {
	Number      int
	Fragment    string
	RoundStates []RoundState
	Dictionary  *Dictionary.Dictionary
}

type RoundState struct {
	Letter   string
	PlayerID int
}

// DidLose determines if round has ended
func (round *Round) LastPlayer() int {

	lengthRoundStates := len(round.RoundStates) - 1
	lastRoundState := round.RoundStates[lengthRoundStates]
	return lastRoundState.PlayerID
}

// LastMove determines if round has ended
func (round *Round) LastMove() RoundState {

	lengthRoundStates := len(round.RoundStates) - 1
	lastRoundState := round.RoundStates[lengthRoundStates]
	return lastRoundState
}

// DidLose determines if a player took the last and losing turn of the game
func (round *Round) DidLose(PlayerID int) bool {
	gameOver := round.IsOver()
	lastPlayerID := round.LastPlayer()
	isLast := lastPlayerID == PlayerID
	return gameOver && isLast
}

// IsOver determines if round has ended
func (round *Round) IsOver() bool {
	return round.Dictionary.WordTree.FragmentIsWord(round.GameState())
}

// GameState does something
func (round *Round) GameState() string {
	compiledFragment := ""
	for _, RoundState := range round.RoundStates {
		compiledFragment += RoundState.Letter
	}

	return compiledFragment
}

// AppendLetter does something
func (round *Round) AppendLetter(letter string, PlayerID int) {
	round.RoundStates = append(round.RoundStates, RoundState{Letter: letter, PlayerID: PlayerID})
}
