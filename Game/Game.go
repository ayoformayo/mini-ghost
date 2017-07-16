package Game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ayoformayo/mini-ghost/Player"
	"github.com/ayoformayo/mini-ghost/Round"
)

// Game stuff
type Game struct {
	Name    string
	Rounds  []Round.Round
	Players []Player.Player
	// letter
}

// StartGame does something
func (p *Game) StartGame() {
	fmt.Println("Hi, my name is", p.Name)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
