package tree_test

import (
	"testing"

	tree "github.com/ayoformayo/mini-ghost/Tree"
)

const targetTestVersion = 1

var trueDictionary = []string{
	"AAHED",
	"AAHING",
	"ZIBELINE",
	"ZIBELLINE",
}

var fakeDictionary = []string{
	"AAHD",
	"AAZING",
	"PPIBELINE",
	"FOOBAR",
}

var wordTree = tree.BuildWordTree(trueDictionary)

func TestFragmentIsWord(t *testing.T) {
	for _, word := range trueDictionary {
		if !wordTree.FragmentIsWord(word) {
			t.Errorf("Tree didn't recognize %s as a word", word)
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
	for _, word := range trueDictionary {
		truncated := word[:3]
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
