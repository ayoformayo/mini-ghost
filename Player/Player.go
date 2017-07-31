package player

import (
	"bufio"
	"math/rand"
	"strings"

	"github.com/ayoformayo/mini-ghost/Dictionary"
	"github.com/ayoformayo/mini-ghost/Round"
)

// Player stuff
type Player struct {
	Dictionary  *Dictionary.Dictionary
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

// Score does b
func (player *Player) Score(round Round.Round) int {
	if round.DidLose(player.ID) {
		return -10
	} else if round.IsOver() {
		return 10
	}
	return 0
}

func (player *Player) minimax(round Round.Round) int {
	if round.IsOver() {
		return player.Score(round)
	}

	scores := []int{}
	moves := []string{}
	options := player.Dictionary.WordTree.GetFragmentChildren(round.GameState())
	for letter, _ := range options {
		newRound := round
		// THIS MUST GET FIXED AND NON HARDCODED
		playerID := 0
		if len(newRound.Moves)%2 != 0 {
			playerID = 1
		}

		newMove := Round.Move{Letter: letter, PlayerID: playerID}
		newRound.Moves = append(newRound.Moves, newMove)
		minimaxed := player.minimax(newRound)
		scores = append(scores, minimaxed)
		moves = append(moves, letter)
	}

	if round.LastPlayer() != player.ID {
		maxScore := -10
		moveIndex := -1
		for i, score := range scores {
			if score >= maxScore {
				maxScore = score
				moveIndex = i
			}
		}
		*player.choice = moves[moveIndex]
		return maxScore
	} else {
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
}

func (player *Player) findAnswer(round Round.Round) string {
	if len(round.GameState()) == 0 {
		return string(letters[rand.Intn(len(letters))])
	}
	player.minimax(round)
	return *player.choice
}

// TakeTurn does something
func (player *Player) TakeTurn(round Round.Round) string {
	var nextLetter string
	if player.IsAI == true {
		player.choice = &nextLetter
		player.findAnswer(round)
		return nextLetter
	} else {
		nextLetter, _ = player.Reader.ReadString('\n')
	}
	if len(nextLetter) < 1 {
		return "1"
	}
	return strings.ToUpper(string(nextLetter[0]))
}
