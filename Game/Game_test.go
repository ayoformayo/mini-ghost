package game_test

import (
	"testing"

	game "github.com/ayoformayo/mini-ghost/Game"
	player "github.com/ayoformayo/mini-ghost/Player"
)

func TestFindPlayer(t *testing.T) {
	var players = []player.Player{
		{ID: 0, IsAI: true},
		{ID: 1, IsAI: true},
		// {ID: 2, IsAI: true, Dictionary: dictionary},
		// {ID: 3, IsAI: true, Dictionary: dictionary},
		// {ID: 4, IsAI: true, Dictionary: dictionary},
	}
	game := game.Game{Players: players}
	foundPlayer := game.FindPlayer(players[0].ID)
	foundPlayer.Letters += "G"
	if game.Players[0].Letters != "G" {
		t.Fatalf("Player %d should have Letter G but has %s", game.Players[0].ID, game.Players[0].Letters)
	}
}

func TestGetLastPlayer(t *testing.T) {
	var players = []player.Player{
		{ID: 0, IsAI: true},
		{ID: 1, IsAI: true},
		// {ID: 2, IsAI: true, Dictionary: dictionary},
		// {ID: 3, IsAI: true, Dictionary: dictionary},
		// {ID: 4, IsAI: true, Dictionary: dictionary},
	}
	game := game.Game{Players: players}
	lastPlayer := game.GetLastPlayer()
	lastPlayer.Letters = "G"
	var playerIndex int
	for i, player := range game.Players {
		if lastPlayer.ID == player.ID {
			playerIndex = i
		}
	}

	actualPlayer := game.Players[playerIndex]

	if actualPlayer.Letters != "G" {
		t.Fatalf("Player %d should have Letter G but has %s", actualPlayer.ID, actualPlayer.Letters)
	}
}
