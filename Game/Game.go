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
		return game.getPlayerCount()
	}
	return number
}

func (game *Game) populatePlayers(count int) {
	i := 0
	for i < count {
		game.Players = append(game.Players,
			Player.Player{Name: fmt.Sprintf("Player %d", i+1),
				Dictionary: game.Dictionary,
				Number:     i,
				IsAI:       i != 0,
				Reader:     game.reader})
		i++
	}
}

func (game *Game) playRound() {
	fmt.Print("What will the first letter be?\n")
	round := Round.Round{Number: len(game.Rounds) + 1, Fragment: ""}
	var lastPlayerNumber int
	for !game.Dictionary.FragmentIsWord(round.Fragment) {
		activePlayer := game.Players[game.ActivePlayer]
		// this needs to be thought out better
		if game.ActivePlayer < len(game.Players)-1 {
			game.ActivePlayer++
		} else {
			game.ActivePlayer = 0
		}
		fmt.Println(fmt.Sprintf("It is %s's turn", activePlayer.Name))
		letter := activePlayer.TakeTurn(round.Fragment)
		round.Fragment += letter
		// to do - clean up if loop and dictionary loop up
		if letter != "1" {
			fmt.Println(fmt.Sprintf("%s wrote %s", activePlayer.Name, letter))
			fmt.Println("")
			fmt.Println(fmt.Sprintf("Phrase is now at %s", round.Fragment))
			lastPlayerNumber = activePlayer.Number
		} else {
			fmt.Println(fmt.Sprintf("%s challenges", activePlayer.Name))
			isEligibleFragment := game.Dictionary.FindEligibleFragment(round.Fragment)
			if len(isEligibleFragment) > 0 || isEligibleFragment == "" {
				fmt.Println("Challenge Successful")
				break
			} else {
				fmt.Println("Challenge Failed")
				lastPlayerNumber = activePlayer.Number
				break
			}
		}
	}

	game.ActivePlayer = lastPlayerNumber
	lastPlayer := &game.Players[game.ActivePlayer]
	ghostLetter := string("GHOST"[len(lastPlayer.Letters)])
	lastPlayer.Letters += ghostLetter
	game.Rounds = append(game.Rounds, round)

	if lastPlayer.Letters != "GHOST" {
		fmt.Println(fmt.Sprintf("%s now has %s", lastPlayer.Name, lastPlayer.Letters))
		fmt.Println()
		game.playRound()
	}
}

// StartGame does something
func (game *Game) StartGame() {
	game.reader = bufio.NewReader(os.Stdin)
	fmt.Print("Welcome to Ghost! Please select number of players - max 5\n")
	playerCount := game.getPlayerCount()
	game.populatePlayers(playerCount)
	game.playRound()
}
