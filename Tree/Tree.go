package tree
import (
  "fmt"
  "regexp"
  "errors"
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

// FragmentIsWord sees if this exists
func (tree *WordTree) FragmentIsWord(fragment string) bool {
	if len(fragment) < 1 {
		return tree.FinalWord == "true"
	}

	asString := string(fragment[:1])
	remainder := string(fragment[1:])

	if _, ok := tree.Letters[asString]; ok {
		return tree.Letters[asString].FragmentIsWord(remainder)
	}
	return false
}



// GetFragmentChildren  sees if this exists
func (tree *WordTree) GetFragmentChildren(fragment string) (map[string]*WordTree, error) {
	if len(fragment) < 1 {
		return tree.Letters, nil
	}

	asString := string(fragment[:1])
	remainder := string(fragment[1:])
  if _, ok := tree.Letters[asString]; !ok{
    return nil, errors.New("This is an invalid phrase")
  }

	return tree.Letters[asString].GetFragmentChildren(remainder)
}

// IsEligible  sees if this exists
func (tree *WordTree) IsEligible(fragment string) bool {
	if len(fragment) < 1 {
		return true
	}

	asString := string(fragment[:1])
	remainder := string(fragment[1:])

	if _, ok := tree.Letters[asString]; ok {
		return tree.Letters[asString].IsEligible(remainder)
	}
	return false
}

// BuildBranches for tree
func (tree *WordTree) BuildBranches(fragment string) {
	if len(fragment) < 1 {
		tree.FinalWord = "true"
		return
	}
	asString := string(fragment[:1])
	remainder := string(fragment[1:])
	if _, ok := tree.Letters[asString]; !ok {
		tree.Letters[asString] = &WordTree{Letters: make(map[string]*WordTree)}
	}
	tree.Letters[asString].BuildBranches(remainder)
}
