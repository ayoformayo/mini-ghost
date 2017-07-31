package player_test

import (
	"fmt"
	"testing"

	"github.com/ayoformayo/mini-ghost/Dictionary"
	player "github.com/ayoformayo/mini-ghost/Player"
	"github.com/ayoformayo/mini-ghost/Round"
)

const targetTestVersion = 1

var fakeDictionary = []string{
	"AAHED",
	"AAHING",
	"ZIBELINE",
	"ZIBELLINE",
}

var wordTree = Dictionary.BuildWordTree(fakeDictionary)

var dictionary = &Dictionary.Dictionary{WordTree: wordTree}

var tests = []player.Player{
	{ID: 0, IsAI: true, Dictionary: dictionary},
	{ID: 1, IsAI: true, Dictionary: dictionary},
	// {ID: 2, IsAI: true, Dictionary: dictionary},
	// {ID: 3, IsAI: true, Dictionary: dictionary},
	// {ID: 4, IsAI: true, Dictionary: dictionary},
	// {"", 0},
	// {" \t\n", 0},
	// {"a", 1},
	// {"f", 4},
	// {"street", 6},
	// {"quirky", 22},
	// {"OXYPHENBUTAZONE", 41},
	// {"alacrity", 13},
}

func TestScoreUnfinished(t *testing.T) {
	roundStates := []Round.Move{
		{Letter: "A", PlayerID: 0},
		{Letter: "A", PlayerID: 1},
		{Letter: "H", PlayerID: 2},
		{Letter: "E", PlayerID: 3},
	}
	round := Round.Round{Moves: roundStates, Dictionary: dictionary}
	for _, test := range tests {
		if actual := test.Score(round); actual != 0 {
			t.Errorf("Player(%q) expected %d, Actual %d", test.ID, 0, actual)
		}
	}

	if player.TestVersion != targetTestVersion {
		t.Fatalf("Found player.TestVersion = %v, want %v.", player.TestVersion, targetTestVersion)
	}
}

func TestScoreFinished(t *testing.T) {
	roundStates := []Round.Move{
		{Letter: "A", PlayerID: 0},
		{Letter: "A", PlayerID: 1},
		{Letter: "H", PlayerID: 0},
		{Letter: "E", PlayerID: 1},
		{Letter: "D", PlayerID: 0},
	}

	round := Round.Round{Moves: roundStates, Dictionary: dictionary}
	losingID := 0
	for _, test := range tests {
		expectedScore := 10
		if test.ID == losingID {
			expectedScore = -10
		}
		if actual := test.Score(round); actual != expectedScore {
			t.Errorf("Player(%q) expected %d, Actual %d", test.ID, expectedScore, actual)
		}
	}

	if player.TestVersion != targetTestVersion {
		t.Fatalf("Found player.TestVersion = %v, want %v.", player.TestVersion, targetTestVersion)
	}
}

func TestTakeTurn(t *testing.T) {
	roundStates := []Round.Move{
		{Letter: "A", PlayerID: 0},
		{Letter: "A", PlayerID: 1},
		{Letter: "H", PlayerID: 0},
		// {Letter: "E", PlayerID: 4},
		// {Letter: "D", PlayerID: 5},
	}
	round := Round.Round{Moves: roundStates, Dictionary: dictionary}
	validLetters := []string{"I", "E"}
	for _, test := range tests {
		isValid := false
		foundAnswer := test.TakeTurn(round)
		for _, letter := range validLetters {
			if letter == foundAnswer {
				isValid = true
			}
		}
		if !isValid {
			t.Errorf("Player(%q)didnt use I or E, used %s", test.ID, isValid, foundAnswer)
		}
	}

	if player.TestVersion != targetTestVersion {
		t.Fatalf("Found player.TestVersion = %v, want %v.", player.TestVersion, targetTestVersion)
	}
}

var oneOnOne = []player.Player{
	{ID: 0, IsAI: true, Dictionary: dictionary},
	{ID: 1, IsAI: true, Dictionary: dictionary},
}

// func TestOneOnOne(t *testing.T) {
// 	roundStates := []Round.Move{
// 		{Letter: "A", PlayerID: 1},
// 		{Letter: "A", PlayerID: 2},
// 		{Letter: "H", PlayerID: 1},
// 		// {Letter: "E", PlayerID: 4},
// 		// {Letter: "D", PlayerID: 5},
// 	}
// 	round := Round.Round{Moves: roundStates, Dictionary: dictionary}
// 	playerTwo := oneOnOne[1]
// 	winningLetter := "E"
// 	for i := 0; i < 100; i++ {
// 		answer := playerTwo.TakeTurn(round)
// 		if answer != winningLetter {
// 			t.Errorf("Player(%q)didnt use E to win, used %s", playerTwo.ID, answer)
// 			break
// 		}
// 	}
//
// 	if player.TestVersion != targetTestVersion {
// 		t.Fatalf("Found player.TestVersion = %v, want %v.", player.TestVersion, targetTestVersion)
// 	}
// }

func TestOneOnOne(t *testing.T) {
	roundStates := []Round.Move{
		{Letter: "Z", PlayerID: 0},
		{Letter: "I", PlayerID: 1},
		{Letter: "B", PlayerID: 0},
		{Letter: "E", PlayerID: 1},
		{Letter: "L", PlayerID: 0},
	}

	round := Round.Round{Moves: roundStates, Dictionary: dictionary}
	playerTwo := oneOnOne[1]
	winningLetter := "L"
	for i := 0; i < 100; i++ {
		answer := playerTwo.TakeTurn(round)
		if answer != winningLetter {
			t.Errorf("Player(%q)didnt use L to win, used %s", playerTwo.ID, answer)
			break
		}
	}

	if player.TestVersion != targetTestVersion {
		t.Fatalf("Found player.TestVersion = %v, want %v.", player.TestVersion, targetTestVersion)
	}
}

func BenchmarkScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range oneOnOne {
			fmt.Println(test)
			// Score(test.input)
		}
	}
}
