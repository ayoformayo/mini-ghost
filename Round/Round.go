package round

import (
	dictionary "github.com/ayoformayo/mini-ghost/Dictionary"
)

// Round is the struct concludes when a player gets a letter
type Round struct {
	Number      int
	Moves       []Move
	PlayerOrder []int
	Dictionary  *dictionary.Dictionary
}

// Move is the struct containing a letter and the player that issued it
type Move struct {
	Letter   string
	PlayerID int
}

// TestVersion sets which tests get fired off
const TestVersion = 1

// IsChallenge returns whether a user issued a letter or challenge
func (move *Move) IsChallenge() bool {
	return move.Letter == "1"
}

// GenerateNextRound creates round for game to pass on down with new player order
func (round *Round) GenerateNextRound() Round {
	if round.IsOver() {
		playerOrder := round.GetNextPlayerOrder()
		round := Round{Dictionary: round.Dictionary, PlayerOrder: playerOrder}
		return round
	}
	return Round{}
}

// GetNextPlayerOrder get list of player order
func (round *Round) GetNextPlayerOrder() []int {
	if len(round.Moves) < 1 {
		return round.PlayerOrder
	}
	var playerIndex int
	for i, playerID := range round.PlayerOrder {
		if round.DidLose(playerID) {
			playerIndex = i
			break
		}
	}

	upToAndIncludingIndex := round.PlayerOrder[:playerIndex]
	afterIndex := round.PlayerOrder[playerIndex:]

	return append(afterIndex, upToAndIncludingIndex...)
}

// LastPlayer returns the last player to have made a move
func (round *Round) LastPlayer() int {

	if len(round.Moves) < 1 {
		return round.PlayerOrder[0]
	}
	lengthMoves := len(round.Moves) - 1
	lastMove := round.Moves[lengthMoves]
	return lastMove.PlayerID
}

// LastMove returns the last move of the round
func (round *Round) LastMove() Move {

	lengthMoves := len(round.Moves) - 1
	var lastMove Move
	if len(round.Moves) > 0 {
		lastMove = round.Moves[lengthMoves]
	}
	return lastMove
}

func (round *Round) didChallenge() bool {
	move := round.LastMove()
	return move.IsChallenge()
}

// DidLose determines if a individual player lost the round
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

// IsOver determines if round has ended, by word completion, challenge or inco
func (round *Round) IsOver() bool {
	lastMove := round.LastMove()
	isChallenge := lastMove.IsChallenge()
	isEliblePhrase := round.Dictionary.WordTree.IsEligible(round.GameState())
	return !isEliblePhrase || round.Dictionary.WordTree.FragmentIsWord(round.GameState()) || isChallenge
}

// GameState returns the phrase assembled in the round thus far
func (round *Round) GameState() string {
	compiledFragment := ""
	for _, Move := range round.Moves {
		compiledFragment += Move.Letter
	}

	return compiledFragment
}

// AppendLetter adds a move
func (round *Round) AppendLetter(letter string, PlayerID int) {
	round.Moves = append(round.Moves, Move{Letter: letter, PlayerID: PlayerID})
}
