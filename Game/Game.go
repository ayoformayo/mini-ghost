package Game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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

func getPlayerCount(reader *bufio.Reader) int {
	playerNumber, _ := reader.ReadString('\n')
	strippedNumber := strings.TrimSpace(playerNumber)
	number, err := strconv.Atoi(strippedNumber)
	if err != nil {
		fmt.Println("Whoops! Not a valid number. Try again")
		getPlayerCount(reader)
	}
	return number
}

func (game *Game) populatePlayers(count int) {
	i := 0
	for i < count {
		game.Players = append(game.Players, Player.Player{})
		i++
	}
}

// StartGame does something
func (game *Game) StartGame() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Welcome to Ghost! Please select number of players - max 5\n")
	playerCount := getPlayerCount(reader)
	game.populatePlayers(playerCount)
	fmt.Println(game.Players)
	fmt.Print("What will the first letter be?\n")
	firstLetter, _ := reader.ReadString('\n')
	fmt.Println(firstLetter)
}
