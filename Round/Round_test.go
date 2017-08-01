package round_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ayoformayo/mini-ghost/Dictionary"
	player "github.com/ayoformayo/mini-ghost/Player"
	round "github.com/ayoformayo/mini-ghost/Round"
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

type testCase struct {
	round    round.Round
	expected []int
}

var tests = []testCase{
	{
		round: round.Round{
			Dictionary:  dictionary,
			PlayerOrder: []int{0, 1},
		},
		expected: []int{1, 0},
	},
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
		test.round.Dictionary = dictionary
		// if test.expected := test.round.GenerateRound(); actual != 0 {
		// 	t.Errorf("Player(%q) expected %d, Actual %d", test.ID, 0, actual)
		// }
		nextRound := test.round.GenerateNextRound()
		if !reflect.DeepEqual(nextRound.PlayerOrder, test.expected) {
			t.Errorf("Round generated %v, not %v", nextRound.PlayerOrder, []int{1, 0})
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
