package game

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
	Dictionary   dictionary.Dictionary
	reader       *bufio.Reader
	// letter
}

// StartGame provides the entry point to the game
func (game *Game) StartGame() {
	game.reader = bufio.NewReader(os.Stdin)
	fmt.Print("Welcome to Ghost! Please select number of players - max 5\n")
	game.GenerateFirstRound()
	game.playRound()
}

// GenerateFirstRound sets up the initial round of the game with players
func (game *Game) GenerateFirstRound() {
	playerCount := game.getPlayerCount()
	game.populatePlayers(playerCount)
	var playerIDs []int
	for _, player := range game.Players {
		playerIDs = append(playerIDs, player.ID)
	}
	firstRound := round.Round{Number: len(game.Rounds) + 1, Dictionary: &game.Dictionary, PlayerOrder: playerIDs}
	game.Rounds = append(game.Rounds, firstRound)
}

// FindPlayer returns a pointer to a Player in the Players Slice
func (game *Game) FindPlayer(playerID int) *player.Player {
	numberOfRounds := len(game.Rounds)
	var playerToFInd player.Player
	if len(game.Players) < 1 {
		return &playerToFInd
	}
	if numberOfRounds < 1 {
		return &game.Players[0]
	}

	for i, player := range game.Players {
		if playerID == player.ID {
			return &game.Players[i]
		}
	}
	return &game.Players[0]
}

// GetLastPlayer finds the last player of the last round and returns a pointer to it on the game
func (game *Game) GetLastPlayer() *player.Player {
	numberOfRounds := len(game.Rounds)
	var lastPlayer *player.Player
	if len(game.Players) < 1 {
		return lastPlayer
	}
	if numberOfRounds < 1 {
		return &game.Players[0]
	}

	lastRound := game.getActiveRound()
	lastPlayerID := lastRound.LastPlayer()
	lastPlayer = game.FindPlayer(lastPlayerID)
	return lastPlayer
}

// Unexported

func (game *Game) getPlayerCount() int {
	playerNumber, _ := game.reader.ReadString('\n')
	strippedNumber := strings.TrimSpace(playerNumber)
	number, err := strconv.Atoi(strippedNumber)
	if err != nil || number < 1 {
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

func (game *Game) getActiveRound() *round.Round {
	numberOfRounds := len(game.Rounds)
	var activeRound *round.Round
	if numberOfRounds < 1 {
		return activeRound
	}

	return &game.Rounds[numberOfRounds-1]
}

func (game *Game) getActivePlayer() *player.Player {
	numberOfRounds := len(game.Rounds)
	var activePlayer *player.Player
	if len(game.Players) < 1 {
		return activePlayer
	}
	if numberOfRounds < 1 {
		return &game.Players[0]
	}

	activeRound := game.getActiveRound()
	playerOrder := activeRound.PlayerOrder
	activeIndex := 0
	if len(activeRound.PlayerOrder) > 0 {
		activeIndex = len(activeRound.Moves) % len(activeRound.PlayerOrder)
	}
	activePlayerID := playerOrder[activeIndex]
	activePlayer = game.FindPlayer(activePlayerID)
	return activePlayer
}

func (game *Game) playRound() {
	round := game.getActiveRound()

	for !round.IsOver() {
		activePlayer := game.getActivePlayer()
		fmt.Println(fmt.Sprintf("It is %s's turn", activePlayer.Name))
		fmt.Print("What will the next letter be?\n")
		letter := activePlayer.TakeTurn(round)
		round.AppendLetter(letter, activePlayer.ID)
		if letter != "1" {
			fmt.Println(fmt.Sprintf("%s wrote %s", activePlayer.Name, letter))
			fmt.Println(fmt.Sprintf("Phrase is now at %s", round.GameState()))
		} else {
			fmt.Println(fmt.Sprintf("%s challenges", activePlayer.Name))
			isEligibleFragment := game.Dictionary.WordTree.IsEligible(round.GameState())
			if isEligibleFragment {
				fmt.Println("Challenge Successful")
				break
			} else {
				fmt.Println("Challenge Failed")
				break
			}
		}
	}
	lastPlayer := game.GetLastPlayer()
	ghostLetter := string("GHOST"[len(lastPlayer.Letters)])
	lastPlayer.Letters += ghostLetter

	if lastPlayer.Letters != "GHOST" {
		fmt.Println(fmt.Sprintf("%s now has %s", lastPlayer.Name, lastPlayer.Letters))
		fmt.Println()
		nextRound := round.GenerateNextRound()
		game.Rounds = append(game.Rounds, nextRound)
		game.playRound()
	} else {
		fmt.Println(fmt.Sprintf("%s has lost!", lastPlayer.Name))
		for _, player := range game.Players {
			fmt.Println(fmt.Sprintf("%s has %s!", player.Name, player.Letters))
		}
	}
}
