package tree
// WordTree is Recursive data structure for branches of word search
type WordTree struct {
	FinalWord string
	Letters   map[string]*WordTree
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
func (tree *WordTree) GetFragmentChildren(fragment string) map[string]*WordTree {
	if len(fragment) < 1 {
		return tree.Letters
	}

	asString := string(fragment[:1])
	remainder := string(fragment[1:])

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
