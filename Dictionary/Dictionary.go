package Dictionary

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// Dictionary struct
type Dictionary struct {
	TotalWords []string
	WordTree   WordTree
}

// LoadEligibleWords gets it in memory
func (dictionary *Dictionary) LoadEligibleWords() {
	lines, _ := ReadLines("Dictionary/EligibleDictionary.txt")
	dictionary.TotalWords = lines
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

// GetFragmentChildren  sees if this exists
func (tree *WordTree) GetFragmentChildren(fragment string) map[string]*WordTree {
	if len(fragment) < 1 {
		return tree.Letters
	}

	asString := string(fragment[:1])
	remainder := string(fragment[1:])

	return tree.Letters[asString].GetFragmentChildren(remainder)
}

// // ResetDictionary sets eligiblewords to all words again
// func (dictionary *Dictionary) ResetDictionary() {
// 	dictionary.EligibleWords = dictionary.TotalWords
// }

// ReadLines opens dictionary file
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// WordTree is Recursive data structure for branches of word search
type WordTree struct {
	FinalWord string
	Letters   map[string]*WordTree
}

func (tree *WordTree) buildBranches(fragment string) {
	if len(fragment) < 1 {
		tree.FinalWord = "true"
		return
	}
	asString := string(fragment[:1])
	remainder := string(fragment[1:])
	if _, ok := tree.Letters[asString]; !ok {
		tree.Letters[asString] = &WordTree{Letters: make(map[string]*WordTree)}
	}
	tree.Letters[asString].buildBranches(remainder)
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
				wordTree.buildBranches(v)
			}
		}
	}
	fmt.Println(wordTree)
	return wordTree
}

func writeEligibleDictionary(lines []string) {
	f, _ := os.Create("Dictionary/EligibleDictionary.txt")
	// check(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, word := range lines {
		withNewline := fmt.Sprintf("%s\n", word)
		writtenWord, _ := w.WriteString(withNewline)
		// w.WriteString("\n")
		fmt.Printf("wrote %d bytes\n", writtenWord)

	}
	w.Flush()
}
