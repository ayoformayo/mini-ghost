package Player

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/ayoformayo/mini-ghost/Dictionary"
)

// Player stuff
type Player struct {
	Dictionary Dictionary.Dictionary
	Letters    string
	Name       string
	Number     int
	IsAI       bool
	Reader     *bufio.Reader
}

// func isLetter(s string) bool {
// 	for _, r := range s {
// 		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
// 			return false
// 		}
// 	}
// 	return true
// }

func (player *Player) findAnswer(fragment string) string {
	return player.Dictionary.FindEligibleFragment(fragment)
}

// RequestLetter does something
func (player *Player) RequestLetter(fragment string) string {
	fmt.Print("Add a valid letter.\n")
	var nextLetter string
	if player.IsAI == true {
		nextLetter = player.findAnswer(fragment)
	} else {
		nextLetter, _ = player.Reader.ReadString('\n')
	}
	return strings.ToUpper(string(nextLetter[0]))
}
