package tree

import (
	"errors"
	"fmt"
	"regexp"
)

// WordTree is Recursive data structure for branches of word search
type WordTree struct {
	FinalWord string
	Letters   map[string]*WordTree
}

// BuildWordTree creates the word tree
func BuildWordTree(lines []string) WordTree {
	preceedingWord := ""
	wordTree := WordTree{Letters: make(map[string]*WordTree)}
	for _, v := range lines {
		if len(v) > 3 {
			reg := fmt.Sprintf("^%s", preceedingWord)
			match, _ := regexp.MatchString(reg, v)
			if len(preceedingWord) == 0 || !match {
				preceedingWord = v
				wordTree.BuildBranches(v)
			}
		}
	}
	return wordTree
}

// HOW CAN WE COLLAPSE FOLLOWING LOGIC

func splitFragment(fragment string) (string, string) {
	return string(fragment[:1]), string(fragment[1:])
}

// FragmentIsWord sees if this exists
func (tree *WordTree) FragmentIsWord(fragment string) bool {
	if len(fragment) < 1 {
		return tree.FinalWord == "true"
	}

	asString, remainder := splitFragment(fragment)
	if _, ok := tree.Letters[asString]; ok {
		return tree.Letters[asString].FragmentIsWord(remainder)
	}

	return false
}

// GetFragmentChildren  get possible options given a certain fragment
func (tree *WordTree) GetFragmentChildren(fragment string) (map[string]*WordTree, error) {
	if len(fragment) < 1 {
		return tree.Letters, nil
	}

	asString, remainder := splitFragment(fragment)
	if _, ok := tree.Letters[asString]; !ok {
		return nil, errors.New("This is an invalid phrase")
	}

	return tree.Letters[asString].GetFragmentChildren(remainder)
}

// IsEligible tests if a phrase has possible descendants
func (tree *WordTree) IsEligible(fragment string) bool {
	if len(fragment) < 1 {
		return true
	}

	asString, remainder := splitFragment(fragment)
	if _, ok := tree.Letters[asString]; ok {
		return tree.Letters[asString].IsEligible(remainder)
	}

	return false
}

// BuildBranches creates the all the word branches for the tree
func (tree *WordTree) BuildBranches(fragment string) {
	if len(fragment) < 1 {
		tree.FinalWord = "true"
		return
	}

	asString, remainder := splitFragment(fragment)
	if _, ok := tree.Letters[asString]; !ok {
		tree.Letters[asString] = &WordTree{Letters: make(map[string]*WordTree)}
	}

	tree.Letters[asString].BuildBranches(remainder)
}
