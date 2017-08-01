package Game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ayoformayo/mini-ghost/Dictionary"
	"github.com/ayoformayo/mini-ghost/Player"
	round "github.com/ayoformayo/mini-ghost/Round"
)

// Game stuff
type Game struct {
	ActivePlayer int
	Name         string
	Rounds       []round.Round
	Players      []player.Player
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
			player.Player{Name: fmt.Sprintf("Player %d", i+1),
				Dictionary: &game.Dictionary,
				ID:         i,
				// IsAI:       true,
				IsAI:        i != 0,
				PlayerCount: count,
				Reader:      game.reader})
		i++
	}
}

func (game *Game) generateFirstRound() round.Round {
	playerCount := game.getPlayerCount()
	game.populatePlayers(playerCount)
	var playerIDs []int
	for _, player := range game.Players {
		playerIDs = append(playerIDs, player.ID)
	}
	return round.Round{Number: len(game.Rounds) + 1, Dictionary: &game.Dictionary, PlayerOrder: playerIDs}
}

func (game *Game) playRound(round round.Round) {
	fmt.Print("What will the next letter be?\n")
	// return round.Round{Number: len(game.Rounds) + 1, Dictionary: &game.Dictionary}
	var lastPlayerID int

	for !round.IsOver() {
		activePlayer := game.Players[game.ActivePlayer]
		// this needs to be thought out better
		if game.ActivePlayer < len(game.Players)-1 {
			game.ActivePlayer++
		} else {
			game.ActivePlayer = 0
		}
		fmt.Println(fmt.Sprintf("It is %s's turn", activePlayer.Name))
		letter := activePlayer.TakeTurn(round)
		round.AppendLetter(letter, activePlayer.ID)
		// to do - clean up if loop and dictionary loop up
		if letter != "1" {
			fmt.Println(fmt.Sprintf("%s wrote %s", activePlayer.Name, letter))
			fmt.Println(fmt.Sprintf("Phrase is now at %s", round.GameState()))
			lastPlayerID = activePlayer.ID
		} else {
			fmt.Println(fmt.Sprintf("%s challenges", activePlayer.Name))
			isEligibleFragment := game.Dictionary.WordTree.IsEligible(round.GameState())
			if isEligibleFragment {
				fmt.Println("Challenge Successful")
				break
			} else {
				fmt.Println("Challenge Failed")
				lastPlayerID = activePlayer.ID
				break
			}
		}
	}

	game.ActivePlayer = lastPlayerID
	lastPlayer := &game.Players[game.ActivePlayer]
	ghostLetter := string("GHOST"[len(lastPlayer.Letters)])
	lastPlayer.Letters += ghostLetter
	game.Rounds = append(game.Rounds, round)

	if lastPlayer.Letters != "GHOST" {
		fmt.Println(fmt.Sprintf("%s now has %s", lastPlayer.Name, lastPlayer.Letters))
		fmt.Println()
		nextRound := round.GenerateNextRound()
		game.playRound(nextRound)
	} else {
		fmt.Println(fmt.Sprintf("%s has lost!", lastPlayer.Name))
		for _, player := range game.Players {
			fmt.Println(fmt.Sprintf("%s has %s!", player.Name, player.Letters))
		}
	}
}

// StartGame does something
func (game *Game) StartGame() {
	game.reader = bufio.NewReader(os.Stdin)
	fmt.Print("Welcome to Ghost! Please select number of players - max 5\n")
	round := game.generateFirstRound()
	game.playRound(round)
}
