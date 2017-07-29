package Dictionary

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// Dictionary struct
type Dictionary struct {
	EligibleWords []string
	TotalWords    []string
}

// LoadEligibleWords gets it in memory
func (dictionary *Dictionary) LoadEligibleWords() {
	lines, _ := ReadLines("Dictionary/EligibleDictionary.txt")
	dictionary.EligibleWords = lines
	dictionary.TotalWords = lines
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// FindEligibleFragment finds next possible letter
func (dictionary *Dictionary) FindEligibleFragment(fragment string) string {
	reg := fmt.Sprintf("^%s.+", fragment)
	letterPosition := len(fragment)
	newDictionary := []string{}
	newFragment := ""
	for _, word := range dictionary.EligibleWords {
		match, _ := regexp.MatchString(reg, word)
		if match {
			if newFragment == "" {
				newFragment = string(word[letterPosition])
			}
			newDictionary = append(newDictionary, word)
		}
	}
	dictionary.EligibleWords = newDictionary
	return newFragment
}

// FragmentIsWord sees if this exists
func (dictionary *Dictionary) FragmentIsWord(fragment string) bool {
	return contains(dictionary.EligibleWords, fragment)
}

// ResetDictionary sets eligiblewords to all words again
func (dictionary *Dictionary) ResetDictionary() {
	dictionary.EligibleWords = dictionary.TotalWords
}

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

// FilterWords chooses only winning words
func FilterWords(lines []string) []string {
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
	return []string{}
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
