package round

import (
	dictionary "github.com/ayoformayo/mini-ghost/Dictionary"
)

// Round stuff
type Round struct {
	Number      int
	Moves       []Move
	PlayerOrder []int
	Dictionary  *dictionary.Dictionary
}

type Move struct {
	Letter   string
	PlayerID int
}

const TestVersion = 1

func (move *Move) IsChallenge() bool {
	return move.Letter == "1"
}

// GenerateNextRound creates round for game to pass on down with new pkayer order
func (round *Round) GenerateNextRound() Round {
	if round.IsOver() {
		playerOrder := round.GetNextPlayerOrder()
		round := Round{Dictionary: round.Dictionary, PlayerOrder: playerOrder}
		return round
	}
	return Round{}
}

// GetNextPlayerOrder get list of player order
func (thisRound *Round) GetNextPlayerOrder() []int {
	if len(thisRound.Moves) < 1 {
		return thisRound.PlayerOrder
	}
	var playerIndex int
	for i, playerID := range thisRound.PlayerOrder {
		if thisRound.DidLose(playerID) {
			playerIndex = i
			break
		}
	}

	upToAndIncludingIndex := thisRound.PlayerOrder[:playerIndex]
	afterIndex := thisRound.PlayerOrder[playerIndex:]

	return append(afterIndex, upToAndIncludingIndex...)
}

// LastPlayer determines if round has ended
func (round *Round) LastPlayer() int {

	if len(round.Moves) < 1 {
		return round.PlayerOrder[0]
	}
	lengthMoves := len(round.Moves) - 1
	lastMove := round.Moves[lengthMoves]
	return lastMove.PlayerID
}

// LastMove determines if round has ended
func (round *Round) LastMove() Move {

	lengthMoves := len(round.Moves) - 1
	var lastMove Move
	if len(round.Moves) > 0 {
		lastMove = round.Moves[lengthMoves]
	}
	return lastMove
}

// LastMove determines if round has ended
func (round *Round) didChallenge() bool {
	move := round.LastMove()
	return move.IsChallenge()
}

// DidLose determines if a player took the last and losing turn of the game
func (round *Round) DidLose(PlayerID int) bool {
	lastMove := round.LastMove()
	isChallenge := lastMove.IsChallenge()
	if isChallenge {
		unabridged := round.GameState()
		length := len(unabridged)
		challenged := unabridged[:length-1]
		isEliblePhrase := round.Dictionary.WordTree.IsEligible(challenged)
		if isEliblePhrase {
			return lastMove.PlayerID == PlayerID
		}
		return round.Moves[length-1].PlayerID != PlayerID

	}
	gameOver := round.IsOver()
	lastPlayerID := round.LastPlayer()
	isLast := lastPlayerID == PlayerID
	return gameOver && isLast
}

// IsOver determines if round has ended
func (round *Round) IsOver() bool {
	lastMove := round.LastMove()
	isChallenge := lastMove.IsChallenge()
	isEliblePhrase := round.Dictionary.WordTree.IsEligible(round.GameState())
	return !isEliblePhrase || round.Dictionary.WordTree.FragmentIsWord(round.GameState()) || isChallenge
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
