package round

import (
	"github.com/ayoformayo/mini-ghost/Dictionary"
)

// Round stuff
type Round struct {
	Number      int
	Moves       []Move
	PlayerOrder []int
	Dictionary  *Dictionary.Dictionary
}

type Move struct {
	Letter   string
	PlayerID int
}

const TestVersion = 1

// GenerateNextRound creates round for game to pass on down with new pkayer order
func (round *Round) GenerateNextRound() Round {
	if round.IsOver() {
		playerOrder := round.getNextPlayerOrder()
		round := Round{Dictionary: round.Dictionary, PlayerOrder: playerOrder}
		return round
	}
	return Round{}
}

func (thisRound *Round) getNextPlayerOrder() []int {
	lastPlayerID := thisRound.LastPlayer()
	var playerIndex int
	for i, playerID := range thisRound.PlayerOrder {
		if playerID == lastPlayerID {
			playerIndex = i
			break
		}
	}

	beforeIndex := thisRound.PlayerOrder[:playerIndex]
	afterIndex := thisRound.PlayerOrder[playerIndex:]
	return append(afterIndex, beforeIndex...)
}

// LastPlayer determines if round has ended
func (round *Round) LastPlayer() int {

	lengthMoves := len(round.Moves) - 1
	lastMove := round.Moves[lengthMoves]
	return lastMove.PlayerID
}

// LastMove determines if round has ended
func (round *Round) LastMove() Move {

	lengthMoves := len(round.Moves) - 1
	lastMove := round.Moves[lengthMoves]
	return lastMove
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
	for _, Move := range round.Moves {
		compiledFragment += Move.Letter
	}

	return compiledFragment
}

// AppendLetter does something
func (round *Round) AppendLetter(letter string, PlayerID int) {
	round.Moves = append(round.Moves, Move{Letter: letter, PlayerID: PlayerID})
}
