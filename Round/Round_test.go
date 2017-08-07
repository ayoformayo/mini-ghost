package round_test

import (
	"fmt"
	"reflect"
	"testing"

	dictionary "github.com/ayoformayo/mini-ghost/Dictionary"
	player "github.com/ayoformayo/mini-ghost/Player"
	round "github.com/ayoformayo/mini-ghost/Round"
	tree "github.com/ayoformayo/mini-ghost/Tree"
)

const targetTestVersion = 1

var fakeDictionary = []string{
	"AAHED",
	"AAHING",
	"ZIBELINE",
	"ZIBELLINE",
}

var wordTree = tree.BuildWordTree(fakeDictionary)

var thisDictionary = &dictionary.Dictionary{WordTree: wordTree}

type testCase struct {
	round               round.Round
	expectedPlayerOrder []int
	expectedLastPlayer  int
}

var tests = []testCase{
	{
		round: round.Round{
			Dictionary:  thisDictionary,
			PlayerOrder: []int{0, 1},
			Moves: []round.Move{
				{Letter: "A", PlayerID: 0},
				{Letter: "A", PlayerID: 1},
				{Letter: "H", PlayerID: 0},
				{Letter: "I", PlayerID: 1},
				{Letter: "N", PlayerID: 0},
				{Letter: "G", PlayerID: 1},
			},
		},
		expectedPlayerOrder: []int{1, 0},
	},
	// {ID: 2, IsAI: true, Dictionary: thisDictionary},
	// {ID: 3, IsAI: true, Dictionary: thisDictionary},
	// {ID: 4, IsAI: true, Dictionary: thisDictionary},
}

func TestAppendLetter(t *testing.T) {
	for _, test := range tests {
		test.round.AppendLetter("T", 0)
		expectedLastMove := test.round.LastMove()
		testMove := round.Move{Letter: "T", PlayerID: 0}
		if !reflect.DeepEqual(testMove, expectedLastMove) {
			t.Errorf("Round returned appended generated %v, not %v", testMove, expectedLastMove)
		}
	}
}

func TestLastPlayer(t *testing.T) {
	for _, test := range tests {
		lastPlayer := test.round.LastPlayer()
		finalIndex := len(test.round.Moves) - 1
		expectedLastPlayer := test.round.Moves[finalIndex]
		if lastPlayer != expectedLastPlayer.PlayerID {
			t.Errorf("Round returned last player as %d, not %d", lastPlayer, expectedLastPlayer.PlayerID)
		}
	}
}

func TestLastMove(t *testing.T) {
	for _, test := range tests {
		lastMove := test.round.LastMove()
		finalIndex := len(test.round.Moves) - 1
		expectedLastMove := test.round.Moves[finalIndex]
		if !reflect.DeepEqual(lastMove, expectedLastMove) {
			t.Errorf("Round returned last player as %v, not %v", lastMove, expectedLastMove)
		}
	}
}

func TestDidLose(t *testing.T) {
	for _, test := range tests {
		expectedLastPlayerID := test.round.LastPlayer()
		for _, playerID := range test.round.PlayerOrder {
			didLose := test.round.DidLose(playerID)
			if playerID != playerID && didLose {

				t.Errorf("Round returned last player as %d, not %d", didLose, expectedLastPlayerID)
			}
		}
	}
}

func TestGenerateNextRound(t *testing.T) {
	roundStates := []round.Move{
		{Letter: "A", PlayerID: 0},
		{Letter: "A", PlayerID: 1},
		{Letter: "H", PlayerID: 0},
		{Letter: "I", PlayerID: 1},
		{Letter: "N", PlayerID: 0},
		{Letter: "G", PlayerID: 1},
	}

	for _, test := range tests {
		test.round.Moves = roundStates
		test.round.Dictionary = thisDictionary
		nextRound := test.round.GenerateNextRound()
		if !reflect.DeepEqual(nextRound.PlayerOrder, test.expectedPlayerOrder) {
			t.Errorf("Round generated %v, not %v", nextRound.PlayerOrder, []int{1, 0})
		}
	}

	if round.TestVersion != targetTestVersion {
		t.Fatalf("Found player.TestVersion = %v, want %v.", player.TestVersion, targetTestVersion)
	}
}

func TestPlayerLose(t *testing.T) {
	roundStates := []round.Move{
		{Letter: "A", PlayerID: 0},
		{Letter: "A", PlayerID: 1},
		{Letter: "H", PlayerID: 0},
		{Letter: "I", PlayerID: 1},
		{Letter: "N", PlayerID: 0},
		{Letter: "1", PlayerID: 1},
	}

	for _, test := range tests {
		test.round.Moves = roundStates
		test.round.Dictionary = thisDictionary
		oneLost := test.round.DidLose(1)
		zeroLost := test.round.DidLose(0)
		if oneLost != true {
			t.Errorf("PlayerLose thought %d lost", 1)
		}
		if zeroLost == true {
			t.Errorf("PlayerLose didnt think %d lost", 0)
		}
	}
}

func TestGenerateFailChallenge(t *testing.T) {
	roundStates := []round.Move{
		{Letter: "A", PlayerID: 0},
		{Letter: "A", PlayerID: 1},
		{Letter: "H", PlayerID: 0},
		{Letter: "I", PlayerID: 1},
		{Letter: "N", PlayerID: 0},
		{Letter: "1", PlayerID: 1},
	}

	for _, test := range tests {
		test.round.Moves = roundStates
		test.round.Dictionary = thisDictionary
		nextRound := test.round.GenerateNextRound()
		if !reflect.DeepEqual(nextRound.PlayerOrder, []int{1, 0}) {
			t.Errorf("TestGenerateFailChallenge Round generated %v, not %v\n", nextRound.PlayerOrder, []int{1, 0})
		}
	}

	if round.TestVersion != targetTestVersion {
		t.Fatalf("Found player.TestVersion = %v, want %v.", player.TestVersion, targetTestVersion)
	}
}

func TestGenerateChallengeSuccess(t *testing.T) {
	roundStates := []round.Move{
		{Letter: "A", PlayerID: 0},
		{Letter: "A", PlayerID: 1},
		{Letter: "H", PlayerID: 0},
		{Letter: "I", PlayerID: 1},
		{Letter: "Q", PlayerID: 0},
		{Letter: "1", PlayerID: 1},
	}

	for _, test := range tests {
		test.round.Moves = roundStates
		test.round.Dictionary = thisDictionary
		nextRound := test.round.GenerateNextRound()
		if !reflect.DeepEqual(nextRound.PlayerOrder, []int{0, 1}) {
			t.Errorf("TestGenerateChallengeSuccess =Round generated %v, not %v", nextRound.PlayerOrder, []int{0, 1})
		}
	}

	if round.TestVersion != targetTestVersion {
		t.Fatalf("Found player.TestVersion = %v, want %v.", player.TestVersion, targetTestVersion)
	}
}

func BenchmarkScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			fmt.Println(test)
			// Score(test.input)
		}
	}
}
