package Game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ayoformayo/mini-ghost/Dictionary"
	"github.com/ayoformayo/mini-ghost/Player"
	"github.com/ayoformayo/mini-ghost/Round"
)

// Game stuff
type Game struct {
	ActivePlayer int
	Name         string
	Rounds       []Round.Round
	Players      []Player.Player
	Dictionary   Dictionary.Dictionary
	reader       *bufio.Reader
	// letter
}

func (game *Game) getPlayerCount() int {
	playerNumber, _ := game.reader.ReadString('\n')
	strippedNumber := strings.TrimSpace(playerNumber)
	number, err := strconv.Atoi(strippedNumber)
	if err != nil {
		fmt.Println("Whoops! Not a valid number. Try again")
		game.getPlayerCount()
	}
	return number
}

func (game *Game) populatePlayers(count int) {
	i := 0
	for i < count {
		game.Players = append(game.Players, Player.Player{Name: fmt.Sprintf("Player %d", i), Reader: game.reader})
		i++
	}
}

func (game *Game) playRound() {
	fmt.Print("What will the first letter be?\n")
	activePlayer := game.Players[game.ActivePlayer]
	firstLetter := activePlayer.RequestLetter()
	round := Round.Round{Number: len(game.Rounds) + 1, Fragment: firstLetter}
	fmt.Println(fmt.Sprintf("Phrase is now at %s", round.Fragment))
	for !game.Dictionary.FragmentIsWord(round.Fragment) {
		letter := activePlayer.RequestLetter()
		round.Fragment += letter
		fmt.Println(fmt.Sprintf("You wrote %s", letter))
		fmt.Println(fmt.Sprintf("Phrase is now at %s", round.Fragment))
		fmt.Println(fmt.Sprintf("You wrote %s", letter))
	}
	game.Rounds = append(game.Rounds, round)
}

// StartGame does something
func (game *Game) StartGame() {
	game.reader = bufio.NewReader(os.Stdin)
	fmt.Print("Welcome to Ghost! Please select number of players - max 5\n")
	playerCount := game.getPlayerCount()
	game.populatePlayers(playerCount)
	game.playRound()
}
