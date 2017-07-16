package Round

import "github.com/ayoformayo/mini-ghost/Dictionary"

// Round stuff
type Round struct {
	Number   int
	Fragment string
}

// IncompleteFragment does something
func (round *Round) IncompleteFragment(dictionary *Dictionary.Dictionary) bool {
	if len(round.Fragment) < 4 {
		return true
	}

	return dictionary.FragmentIsWord(round.Fragment)
}
