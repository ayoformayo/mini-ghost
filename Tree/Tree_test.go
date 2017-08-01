package tree_test

import (
	"testing"

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

func TestFragmentIsWord(t *testing.T) {
	for _, word := range fakeDictionary {
		if !wordTree.FragmentIsWord(word) {
			t.Errorf("Tree didn't recognize %s as a word", word)
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
