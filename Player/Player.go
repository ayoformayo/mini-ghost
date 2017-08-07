package player

import (
	"bufio"
	"math/rand"
	"strings"

	dictionary "github.com/ayoformayo/mini-ghost/Dictionary"
	"github.com/ayoformayo/mini-ghost/Round"
)

// Player stuff
type Player struct {
	Dictionary  *dictionary.Dictionary
	Letters     string
	Name        string
	ID          int
	PlayerCount int
	IsAI        bool
	choice      *string
	Reader      *bufio.Reader
}

// TestVersion for testing
const TestVersion = 1
const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Score evaluates the end value of any given move
func (player *Player) Score(thisRound round.Round) int {
	if thisRound.DidLose(player.ID) {
		return -10
	} else if thisRound.IsOver() {
		return 10
	}
	return 0
}

// Minimax does the recursive heavy lifting to identify optimal move choice
func (player *Player) Minimax(thisRound round.Round) int {
	if thisRound.IsOver() {
		return player.Score(thisRound)
	}

	scores := []int{}
	moves := []string{}
	options, _ := player.Dictionary.WordTree.GetFragmentChildren(thisRound.GameState())
	for letter := range options {
		newRound := thisRound
		playerID := 0
		if len(thisRound.PlayerOrder) > 0 {
			playerID = len(newRound.Moves) % len(thisRound.PlayerOrder)
		}

		newMove := round.Move{Letter: letter, PlayerID: playerID}
		newRound.Moves = append(newRound.Moves, newMove)
		minimaxed := player.Minimax(newRound)
		scores = append(scores, minimaxed)
		moves = append(moves, letter)
	}

	if thisRound.LastPlayer() != player.ID {
		maxScore := -10
		moveIndex := -0
		for i, score := range scores {
			if score >= maxScore {
				maxScore = score
				moveIndex = i
			}
		}
		if len(moves) > 0 {
			*player.choice = moves[moveIndex]
		}
		return maxScore
	}
	minScore := 10
	moveIndex := 0
	for i, score := range scores {
		if score <= minScore {
			minScore = score
			moveIndex = i
		}
	}
	*player.choice = moves[moveIndex]
	return minScore
}

func (player *Player) findAnswer(thisRound round.Round) string {
	if len(thisRound.GameState()) == 0 {
		return string(letters[rand.Intn(len(letters))])
	}
	player.Minimax(thisRound)
	return *player.choice
}

// TakeTurn either asks a human player to enter a letter or passes the baton to the AI
func (player *Player) TakeTurn(thisRound *round.Round) string {
	var nextLetter string
	if player.IsAI == true {
		player.choice = &nextLetter
		nextLetter = player.findAnswer(*thisRound)
	} else {
		nextLetter, _ = player.Reader.ReadString('\n')
	}
	if len(nextLetter) < 1 {
		return "1"
	}
	return strings.ToUpper(string(nextLetter[0]))
}
