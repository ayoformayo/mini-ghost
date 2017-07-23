package Player

import (
	"bufio"
	"fmt"
	"math/rand"
	"strings"

	"github.com/ayoformayo/mini-ghost/Dictionary"
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

// func isLetter(s string) bool {
// 	for _, r := range s {
// 		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
// 			return false
// 		}
// 	}
// 	return true
// }

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (player *Player) findAnswer(fragment string) string {
	if len(fragment) == 0 {
		return string(letters[rand.Intn(len(letters))])
	}
	return player.Dictionary.FindEligibleFragment(fragment)
}

// TakeTurn does something
func (player *Player) TakeTurn(fragment string) string {
	fmt.Print("Add a valid letter.\n")
	var nextLetter string
	if player.IsAI == true {
		nextLetter = player.findAnswer(fragment)
	} else {
		nextLetter, _ = player.Reader.ReadString('\n')
	}
	if len(nextLetter) < 1 {
		return "1"
	}
	return strings.ToUpper(string(nextLetter[0]))
}
