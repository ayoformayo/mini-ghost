package tree_test

import (
	"reflect"
	"testing"

	tree "github.com/ayoformayo/mini-ghost/Tree"
)

const targetTestVersion = 1

type testCase struct {
	word  string
	moves []string
}

var trueDictionary = []testCase{
	{word: "AAHED", moves: []string{"E", "I"}},
	{word: "AAHING", moves: []string{"E", "I"}},
	{word: "ZIBELINE", moves: []string{"E"}},
	{word: "ZIBELLINE", moves: []string{"E"}},
}

var fakeDictionary = []string{
	"AAHD",
	"AAZING",
	"PPIBELINE",
	"FOOBAR",
	"ZZZZZZ",
	"QWERTY",
}

func dictionaryWords() []string {
	var dictionaryWords = []string{}
	for _, test := range trueDictionary {
		dictionaryWords = append(dictionaryWords, test.word)
	}
	return dictionaryWords
}

var wordTree = tree.BuildWordTree(dictionaryWords())

func TestGetFragmentChildren(t *testing.T) {
	for _, test := range trueDictionary {
		truncated := test.word[:3]
		children, _ := wordTree.GetFragmentChildren(truncated)
		moves := []string{}
		for key := range children {
			moves = append(moves, key)
		}
		if !reflect.DeepEqual(moves, test.moves) {
			t.Errorf("Tree returned %v as possible moves instead of %v", moves, test.moves)
		}
	}
}

// func TestInvalidFragmentChildren(t *testing.T) {
// 	fmt.Println(wordTree.Letters["A"].Letters["A"].Letters["H"])
// 	for _, test := range fakeDictionary {
// 		children, err := wordTree.GetFragmentChildren(test)
// 		fmt.Println(children)
// 		if err != nil {
// 			t.Errorf("Tree didn't raise error getting children on %s, an invalid wor", test)
// 		}
// 	}
// }

func TestFragmentIsWord(t *testing.T) {
	for _, test := range trueDictionary {
		if !wordTree.FragmentIsWord(test.word) {
			t.Errorf("Tree didn't recognize %s as a word", test.word)
		}
	}
}

func TestFragmentIsNotWord(t *testing.T) {
	for _, fake := range fakeDictionary {
		if wordTree.FragmentIsWord(fake) {
			t.Errorf("Tree recognized %s as a playable word", fake)
		}
	}
}

func TestFragmentIsEligible(t *testing.T) {
	for _, test := range trueDictionary {
		truncated := test.word[:3]
		if !wordTree.IsEligible(truncated) {
			t.Errorf("Tree didn't recognize %s as an eligible phrase", truncated)
		}
	}
}

func TestFragmentIsNotEligible(t *testing.T) {
	for _, fake := range fakeDictionary {
		if wordTree.IsEligible(fake) {
			t.Errorf("Tree recognized %s as an legitimate phrase", fake)
		}
	}
}

// func BenchmarkScore(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// for _, test := range oneOnOne {
// 	fmt.Println(test)
// 	// Score(test.input)
// }
// 	}
// }
