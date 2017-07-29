package Player

import (
	"bufio"
	"fmt"
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
	Number      int
	PlayerCount int
	IsAI        bool
	Reader      *bufio.Reader
}

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// func minimax(options []string) int {
//
// }

func (player *Player) findAnswer(round Round.Round) string {
	if len(round.GameState()) == 0 {
		return string(letters[rand.Intn(len(letters))])
	}
	options := player.Dictionary.WordTree.GetFragmentChildren(round.GameState())
	continuingMoves := []string{}
	finishingMoves := []string{}
	for key := range options {
		nextOption := round.GameState() + key
		if !player.Dictionary.WordTree.FragmentIsWord(nextOption) {
			continuingMoves = append(continuingMoves, key)
		} else {
			finishingMoves = append(continuingMoves, key)
		}
	}
	fmt.Println(continuingMoves)
	fmt.Println(finishingMoves)
	if len(continuingMoves) > 0 {
		return continuingMoves[0]
	}
	return finishingMoves[0]
}

// TakeTurn does something
func (player *Player) TakeTurn(round Round.Round) string {
	fmt.Print("Add a valid letter.\n")
	var nextLetter string
	if player.IsAI == true {
		nextLetter = player.findAnswer(round)
	} else {
		nextLetter, _ = player.Reader.ReadString('\n')
	}
	if len(nextLetter) < 1 {
		return "1"
	}
	return strings.ToUpper(string(nextLetter[0]))
}
