package Player

import (
	"bufio"
	"fmt"
)

// Player stuff
type Player struct {
	Letters string
	Name    string
	Reader  *bufio.Reader
}

// func isLetter(s string) bool {
// 	for _, r := range s {
// 		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
// 			return false
// 		}
// 	}
// 	return true
// }

// RequestLetter does something
func (player *Player) RequestLetter() string {
	fmt.Print("Add a valid letter.\n")
	nextLetter, _ := player.Reader.ReadString('\n')
	return string(nextLetter[0])
}
